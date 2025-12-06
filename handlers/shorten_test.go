package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestRedis(t *testing.T) {
	mockRedis, err := miniredis.Run()
	assert.NoError(t, err, "Failed to start mock Redis server")

	t.Cleanup(func() {
		mockRedis.Close()
	})

	redisURL := "redis://localhost:" + mockRedis.Port()
	os.Setenv("REDIS_URL", redisURL)
	os.Setenv("DEFAULT_KEY_LENGTH", "8")
	os.Setenv("PUBLIC_URL", "http://localhost:8080")

	// Explicitly initialize the Store with the test Redis URL
	InitializeStore(redisURL)
}

func TestShortenSuccess(t *testing.T) {
	setupTestRedis(t)

	body := map[string]string{
		"url": "https://example.com",
	}

	b, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/shorten", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	ShortenHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status 200")

	assert.Contains(t, rr.Body.String(), "http", "Response should contain shortened URL")
}

func TestShortenInvalidURL(t *testing.T) {
	setupTestRedis(t)

	body := map[string]string{
		"url": "not-a-url",
	}

	b, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/shorten", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	ShortenHandler(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code, "Expected status 422")
}
