package utils

import (
	"math/rand"
	"strconv"
	"tiiny-go/config"
	"time"
)

var (
	keyLength int
	rng       *rand.Rand
)

func init() {
	cfg := config.GetConfig()
	keyLength = cfg.DefaultKeyLength
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
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
	cfg := config.GetConfig() // Fetch configuration dynamically
	if domain := cfg.PublicURL; domain != "" {
		return domain
	}
	return hostname + ":" + strconv.Itoa(port) // dev fallback
}
