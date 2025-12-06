package handlers

import (
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"WebsiteName":   "Tiiny",
		"WebsiteDomain": os.Getenv("PUBLIC_URL"),
	}

	err := Templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
