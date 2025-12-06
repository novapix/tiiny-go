package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"tiiny-go/config"
	"tiiny-go/store"
	"tiiny-go/utils"
	"tiiny-go/validation"
	"time"
)

var Store store.URLStore

type ShortenRequest struct {
	URL string `json:"url"`
	Key string `json:"key,omitempty"`
}

func InitializeStore(redisURL string) {
	Store = store.NewRedisStore(redisURL)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Validate input
	if err := validation.ValidateShortenRequest(validation.ShortenRequest{
		URL: req.URL,
		Key: req.Key,
	}); err != nil {
		writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if req.Key == "" {
		req.Key = utils.GenerateKey()
	}

	err := Store.Save(req.Key, req.URL)
	if err != nil {
		log.Println("Error saving URL:", err)
		writeJSONError(w, http.StatusInternalServerError, "Failed to save URL")
		return
	}

	domain := config.GetConfig().PublicURL
	port, err := strconv.Atoi(config.GetConfig().Port) // Convert Port to int
	if err != nil {
		log.Println("Invalid port in configuration:", err)
		writeJSONError(w, http.StatusInternalServerError, "Server configuration error")
		return
	}

	resp := map[string]string{
		"short_url": utils.GenerateDomainName(domain, port) + "/" + req.Key,
		"key":       req.Key,
		"url":       req.URL,
		"expires":   time.Now().Add(30 * 24 * time.Hour).Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:] // Remove the leading "/"

	url, err := Store.Get(key)
	if err != nil || url == "" {
		// If the key is not found, return a 404 error
		writeJSONError(w, http.StatusNotFound, "Short URL not found")
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		HomeHandler(w, r)
		return
	}

	// Otherwise, treat it as a short URL key
	key := strings.TrimPrefix(r.URL.Path, "/")
	url, err := Store.Get(key)
	if err != nil || url == "" {
		writeJSONError(w, http.StatusNotFound, "Short URL not found")
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
