package handlers

import (
	"embed"
	"html/template"
	"log"
)

var Templates *template.Template

func LoadTemplates(fs embed.FS) {
	var err error
	Templates, err = template.ParseFS(fs, "templates/*.html")
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}
}
