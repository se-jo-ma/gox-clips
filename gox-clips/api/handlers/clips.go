package handlers

import (
	"fmt"
	"net/http"
)

func LoadEnvironmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>HTMX Loaded Content: This is dynamically loaded!</p>")
}
