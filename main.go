package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"tiiny-go/config"
	"tiiny-go/handlers"
)

//go:embed static/*
//go:embed templates/*.html
var content embed.FS

func main() {
	cfg := config.GetConfig()

	fmt.Println("Public URL:", cfg.PublicURL)

	handlers.InitializeStore(cfg.RedisURL)

	fs := http.FileServer(http.FS(content))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// handlers.LoadTemplates()
	http.Handle("/static/", fs)

	// Pass the embedded FS to the template loader
	handlers.LoadTemplates(content)

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/shorten", handlers.ShortenHandler)

	log.Printf("Server running at :%s\n", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
