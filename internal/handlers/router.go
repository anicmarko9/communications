package handlers

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"

	"communications/internal/config"
	"communications/internal/utils"
)

// Provides access to the database connection pool and configuration.
// Used to give all route handlers access to .env variables and the DB pool.
type Handler struct {
	Pool *pgxpool.Pool
	Cfg  *config.Config
}

// Sets up the Gin router with middleware (CORS, security headers, compression, body size, rate limiting),
// applies API versioning and route definitions, and returns the configured Gin engine.
// Used to initialize the HTTP server.
func Init(cfg *config.Config, db *pgxpool.Pool) *gin.Engine {
	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	router.Use(setCORS(cfg))
	router.Use(setSecurityHeaders())
	router.Use(setCompression())
	router.Use(setBodySize())
	router.Use(setRateLimiter(cfg))

	handler := &Handler{Pool: db, Cfg: cfg}

	v1 := router.Group("/api/v1")

	v1.GET("/health", handler.HealthHandler)

	v1.POST("/leads/:id", handler.LeadHandler)

	return router
}

// Configures CORS middleware using allowed origins from config.
// Ensures only trusted origins can access the API.
func setCORS(cfg *config.Config) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// Adds basic security headers to all responses.
// Helps protect against common web vulnerabilities.
func setSecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-XSS-Protection", "0")

		c.Next()
	}
}

// Enables GZIP compression for all responses.
// Reduces bandwidth usage for API clients.
func setCompression() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}

// Limits the maximum size of incoming request bodies.
// Prevents abuse by restricting large payloads (up to 25MB, suitable for handling PDF uploads).
func setBodySize() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 25*utils.MB)

		c.Next()
	}
}

var rateLimiterCache = cache.New(time.Minute, 5*time.Minute)

// Returns a rate limiter for the given IP address, creating one if it doesn't exist.
// Used to enforce per-IP rate limiting.
func getRateLimiter(ip string, eventsPerSecond rate.Limit, burst int) *rate.Limiter {
	limiter, found := rateLimiterCache.Get(ip)
	if found {
		return limiter.(*rate.Limiter)
	}

	limiter = rate.NewLimiter(eventsPerSecond, burst)
	rateLimiterCache.Set(ip, limiter, cache.DefaultExpiration)

	return limiter.(*rate.Limiter)
}

// Applies a rate limiter middleware to all requests based on client IP.
// Returns HTTP 429 if the client exceeds the allowed request rate.
func setRateLimiter(cfg *config.Config) gin.HandlerFunc {
	eventsPerSecond := rate.Every(time.Duration(cfg.ThrottleTTL) * time.Second)
	burst := cfg.ThrottleLimit

	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getRateLimiter(ip, eventsPerSecond, burst)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, utils.APIResponse[utils.DefaultResponse]{
				Meta: utils.Meta{
					Status:    utils.StatusError,
					Message:   "Too many requests. Please try again later.",
					Timestamp: utils.GetCurrentTimestamp(),
				},
				Data: utils.DefaultResponse{},
			})

			return
		}

		c.Next()
	}
}
