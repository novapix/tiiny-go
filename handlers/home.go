package handlers

import (
	"html/template"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"WebsiteName":   "Tiiny",
		"WebsiteDomain": os.Getenv("PUBLIC_URL"),
		"Mode":          "light",
	}

	tmpl.Execute(w, data)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Mode": "light",
	}

	w.WriteHeader(http.StatusNotFound)
	tmpl.Execute(w, data)
}
