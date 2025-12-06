package handlers

import (
	"html/template"
	"log"
)

var Templates *template.Template

func LoadTemplates() {
	var err error
	Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}
}
