package store

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

func TestRedisStore(t *testing.T) {
	// Start a mock Redis server
	mockRedis, err := miniredis.Run()
	assert.NoError(t, err, "Failed to start mock Redis server")
	defer mockRedis.Close()

	// Use the mock Redis server's address
	redisURL := "redis://" + mockRedis.Addr()
	store := NewRedisStore(redisURL)

	// Test Save and Get
	key := "test-key"
	url := "https://example.com"

	err = store.Save(key, url)
	assert.NoError(t, err, "Save should not return an error")

	retrievedURL, err := store.Get(key)
	assert.NoError(t, err, "Get should not return an error")
	assert.Equal(t, url, retrievedURL, "Retrieved URL should match the saved URL")

	// Test Get for a non-existent key
	nonExistentKey := "non-existent-key"
	retrievedURL, err = store.Get(nonExistentKey)
	assert.Error(t, err, "Get should return an error for a non-existent key")
	assert.Empty(t, retrievedURL, "Retrieved URL should be empty for a non-existent key")
}
