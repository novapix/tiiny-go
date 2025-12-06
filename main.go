package main

import (
	"fmt"
	"log"
	"net/http"
	"tiiny-go/config"
	"tiiny-go/handlers"
)

func main() {

	cfg := config.LoadConfig()
	fmt.Println("Public URL:", cfg.PublicURL)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/":
			handlers.HomeHandler(w, r)
		case r.URL.Path == "/shorten":
			handlers.ShortenHandler(w, r)
		default:
			handlers.RedirectHandler(w, r)
		}
	})

	// Custom 404 route
	http.HandleFunc("/404", handlers.NotFoundHandler)

	log.Printf("Server running at :%s\n", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
