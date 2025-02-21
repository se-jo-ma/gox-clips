package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("internal/templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoadContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>HTMX Loaded Content: This is dynamically loaded!</p>")
}

