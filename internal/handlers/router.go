package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"communications/internal/config"
	"communications/internal/services"
	"communications/internal/utils"
)

var (
	limiters = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func Init(cfg *config.Config) *gin.Engine {
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

	v1 := router.Group("/api/v1")

	v1.GET("/health", HealthHandler)

	return router
}

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

func setSecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-XSS-Protection", "0")

		c.Next()
	}
}

func setCompression() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}

func setBodySize() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 25*utils.MB)

		c.Next()
	}
}

func getRateLimiter(ip string, eventsPerSecond rate.Limit, burst int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := limiters[ip]

	if !exists {
		limiter = rate.NewLimiter(eventsPerSecond, burst)
		limiters[ip] = limiter
	}

	return limiter
}

func setRateLimiter(cfg *config.Config) gin.HandlerFunc {
	eventsPerSecond := rate.Every(time.Duration(cfg.ThrottleTTL) * time.Second)
	burst := cfg.ThrottleLimit

	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getRateLimiter(ip, eventsPerSecond, burst)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, services.Throttler())
			return
		}

		c.Next()
	}
}
