package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	GinMode string
}

func LoadConfig() *Config {
	godotenv.Load(".env")

	required := []string{
		"PORT",
		"GIN_MODE",
	}

	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Environment variable %s must be set", key)
		}
	}

	return &Config{
		Port:    os.Getenv("PORT"),
		GinMode: os.Getenv("GIN_MODE"),
	}
}
