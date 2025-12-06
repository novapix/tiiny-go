package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"tiiny-go/utils"
	"tiiny-go/validation"
	"time"
)

type ShortenRequest struct {
	URL string `json:"url"`
	Key string `json:"key,omitempty"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if err := validation.ValidateShortenRequest(validation.ShortenRequest{
		URL: req.URL,
		Key: req.Key,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if req.Key == "" {
		req.Key = utils.GenerateKey()
	}

	err := utils.SaveURL(req.Key, req.URL)
	if err != nil {
		log.Println("Error saving URL:", err)
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{
		"short_url": utils.GenerateDomainName(r.Host, 8080) + "/" + req.Key,
		"key":       req.Key,
		"url":       req.URL,
		"expires":   time.Now().Add(30 * 24 * time.Hour).Format(time.RFC3339), // optional
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// RedirectHandler redirects to the original URL
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:] // remove leading "/"

	url, err := utils.GetURL(key)
	if err != nil || url == "" {
		NotFoundHandler(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
