package handlers

import (
	"net/http"

	"github.com/boonyarit-iamsaard/templ-htmx-go/views"
)

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
