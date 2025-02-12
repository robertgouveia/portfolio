package http

import (
	"github.com/robertgouveia/portfolio/internal/project"
	"net/http"
)

func MapRoutes(mux *http.ServeMux, h project.Handler) {
	mux.HandleFunc("/projects", h.GetProjects())
}
