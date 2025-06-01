package config

import (
	"communications/internal/utils"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Holds all environment variables and settings required to run the application.
// Populated once at startup and used throughout the app for configuration.
type Config struct {
	Port             string   // Port the server listens on.
	ThrottleTTL      int      // Time-to-live for request throttling (seconds).
	ThrottleLimit    int      // Maximum requests allowed per TTL.
	GinMode          string   // Gin framework mode (debug, release, etc.).
	AllowedOrigins   []string // List of allowed CORS origins.
	DatabaseHost     string   // PostgreSQL host.
	DatabasePort     int      // PostgreSQL port.
	DatabaseUser     string   // PostgreSQL user.
	DatabasePassword string   // PostgreSQL password.
	DatabaseName     string   // PostgreSQL database name.
	DatabaseSSL      string   // PostgreSQL SSL mode.
	AzureURL         string   // Azure service endpoint.
	EmailFrom        string   // Default sender Email address.
	SMSFrom          string   // Default sender SMS address.
}

// Loads environment variables from .env (if present) or from the environment, validates required variables, and sets server timezone to UTC.
// Called once at startup to initialize application configuration.
// Returns a pointer to the populated Config struct, or terminates the app if required variables are missing.
func Load() *Config {
	godotenv.Load(".env")

	required := []string{
		"PORT",
		"THROTTLE_TTL",
		"THROTTLE_LIMIT",
		"GIN_MODE",
		"ALLOWED_ORIGINS",
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB",
		"POSTGRES_SSL",
		"AZURE_URL",
		"EMAIL_FROM",
		"SMS_FROM",
	}

	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Environment variable %s must be set", key)
		}
	}

	time.Local = time.UTC

	return &Config{
		Port:             os.Getenv("PORT"),
		ThrottleTTL:      utils.StringToNumber[int](os.Getenv("THROTTLE_TTL")),
		ThrottleLimit:    utils.StringToNumber[int](os.Getenv("THROTTLE_LIMIT")),
		GinMode:          os.Getenv("GIN_MODE"),
		AllowedOrigins:   utils.SplitString(os.Getenv("ALLOWED_ORIGINS"), ","),
		DatabaseHost:     os.Getenv("POSTGRES_HOST"),
		DatabasePort:     utils.StringToNumber[int](os.Getenv("POSTGRES_PORT")),
		DatabaseUser:     os.Getenv("POSTGRES_USER"),
		DatabasePassword: os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName:     os.Getenv("POSTGRES_DB"),
		DatabaseSSL:      os.Getenv("POSTGRES_SSL"),
		AzureURL:         os.Getenv("AZURE_URL"),
		EmailFrom:        os.Getenv("EMAIL_FROM"),
		SMSFrom:          os.Getenv("SMS_FROM"),
	}
}
