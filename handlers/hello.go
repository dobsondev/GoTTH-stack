package handlers

import (
	"net/http"

	"github.com/dobsondev/gotth-stack/templ/components"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Get straight HTML from a templ component to return to HTMX
	components.TextBlock("Hello!").Render(r.Context(), w)
}
