package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicURL        string
	DefaultKeyLength int
	RedisURL         string
	Port             string
}

func LoadConfig() *Config {
	err_load := godotenv.Load()
	if err_load != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{}

	cfg.PublicURL = os.Getenv("PUBLIC_URL")
	if cfg.PublicURL == "" {
		cfg.PublicURL = "http://localhost:8080" // fallback for local dev
	}

	// Default key length
	if keyLengthStr := os.Getenv("DEFAULT_KEY_LENGTH"); keyLengthStr != "" {
		if length, err := strconv.Atoi(keyLengthStr); err == nil {
			cfg.DefaultKeyLength = length
		} else {
			cfg.DefaultKeyLength = 8
			log.Println("Invalid DEFAULT_KEY_LENGTH, using 8")
		}
	} else {
		cfg.DefaultKeyLength = 8
	}

	// Redis connection URL
	cfg.RedisURL = os.Getenv("REDIS_URL")
	if cfg.RedisURL == "" {
		cfg.RedisURL = "redis://localhost:6379" // local dev fallback
	}

	// App port
	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg
}
