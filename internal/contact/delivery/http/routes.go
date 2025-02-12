package http

import (
	"github.com/robertgouveia/portfolio/internal/contact"
	"net/http"
)

func MapRoutes(mux *http.ServeMux, h contact.Handler) {
	mux.HandleFunc("/contact", h.GetContact())
}
