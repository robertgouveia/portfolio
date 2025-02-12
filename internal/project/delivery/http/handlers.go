package http

import (
	"github.com/robertgouveia/portfolio/internal/project"
	"net/http"
)

type ProjectHandler struct {
	render func(w http.ResponseWriter, data map[string]interface{}, layout string, page string, partial ...string)
}

func NewProjectHandler(r func(w http.ResponseWriter, data map[string]interface{}, layout string, page string, partial ...string)) project.Handler {
	return &ProjectHandler{
		render: r,
	}
}

func (h *ProjectHandler) GetProjects() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.render(w, nil, "base", "projects", "nav")
	}
}
