package http

import (
	"github.com/robertgouveia/portfolio/internal/contact"
	"net/http"
)

type ContactHandler struct {
	render func(w http.ResponseWriter, data map[string]interface{}, layout string, page string, partial ...string)
}

func NewContactHandler(r func(w http.ResponseWriter, data map[string]interface{}, layout string, page string, partial ...string)) contact.Handler {
	return &ContactHandler{
		render: r,
	}
}

func (h *ContactHandler) GetContact() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.render(w, nil, "base", "contact", "nav")
	}
}
