package main

import (
	"fmt"
	"log"
	"net/http"
	"tiiny-go/config"
	"tiiny-go/handlers"
)

func main() {
	cfg := config.GetConfig()

	fmt.Println("Public URL:", cfg.PublicURL)

	handlers.InitializeStore(cfg.RedisURL)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	handlers.LoadTemplates()

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/shorten", handlers.ShortenHandler)

	log.Printf("Server running at :%s\n", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
