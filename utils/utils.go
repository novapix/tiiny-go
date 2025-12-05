package utils

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	keyLength int
	rng       *rand.Rand // local RNG, not global
)

func init() {
	lengthStr := os.Getenv("DEFAULT_KEY_LENGTH")
	if lengthStr == "" {
		keyLength = 8
	} else {
		l, err := strconv.Atoi(lengthStr)
		if err != nil {
			keyLength = 8
		} else {
			keyLength = l
		}
	}

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
	if domain := os.Getenv("PUBLIC_URL"); domain != "" {
		return domain
	}
	return hostname + ":" + strconv.Itoa(port) // dev fallback
}
