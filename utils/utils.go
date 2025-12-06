package utils

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"tiiny-go/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	keyLength   int
	rng         *rand.Rand
	redisClient *redis.Client
	ctx         = context.Background()
	cfg         *config.Config
)

func init() {
	cfg = config.LoadConfig()
	keyLength = cfg.DefaultKeyLength

	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	opt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		panic(fmt.Sprintf("Invalid REDIS_URL: %v", err))
	}

	redisClient = redis.NewClient(opt)
}

func GenerateKey() string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, keyLength)
	for i := range result {
		result[i] = chars[rng.Intn(len(chars))]
	}
	return string(result)
}

func GenerateDomainName(hostname string, port int) string {
	if domain := cfg.PublicURL; domain != "" {
		return domain
	}
	return hostname + ":" + strconv.Itoa(port) // dev fallback
}

func SaveURL(key string, url string) error {
	ttl := 30 * 24 * time.Hour // 30 days
	return redisClient.Set(ctx, key, url, ttl).Err()
}

func GetURL(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}
