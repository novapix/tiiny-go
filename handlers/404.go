package handlers

import (
	"log"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	err := Templates.ExecuteTemplate(w, "404.html", map[string]interface{}{
		"Mode": "light",
	})

	if err != nil {
		log.Println("404 template error:", err)
	}
}
