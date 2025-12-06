package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once     sync.Once
	instance *Config
)

func Load() *Config {
	_ = godotenv.Load()

	keyLen := 8
	if v := os.Getenv("DEFAULT_KEY_LENGTH"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			keyLen = parsed
		}
	}

	return &Config{
		Port:             getEnv("PORT", "8080"),
		PublicURL:        getEnv("PUBLIC_URL", "http://localhost:8080"),
		DefaultKeyLength: keyLen,
		RedisURL:         getEnv("REDIS_URL", "localhost:6379"),
	}
}

func GetConfig() *Config {
	once.Do(func() {
		instance = Load()
	})
	return instance
}

func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
