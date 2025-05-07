package config

import (
	"communications/internal/utils"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	ThrottleTTL    int
	ThrottleLimit  int
	GinMode        string
	AllowedOrigins []string
}

func Load() *Config {
	godotenv.Load(".env")

	required := []string{
		"PORT",
		"THROTTLE_TTL",
		"THROTTLE_LIMIT",
		"GIN_MODE",
		"ALLOWED_ORIGINS",
	}

	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Environment variable %s must be set", key)
		}
	}

	time.Local = time.UTC

	return &Config{
		Port:           os.Getenv("PORT"),
		ThrottleTTL:    utils.StringToNumber[int](os.Getenv("THROTTLE_TTL")),
		ThrottleLimit:  utils.StringToNumber[int](os.Getenv("THROTTLE_LIMIT")),
		GinMode:        os.Getenv("GIN_MODE"),
		AllowedOrigins: utils.SplitString(os.Getenv("ALLOWED_ORIGINS"), ","),
	}
}
