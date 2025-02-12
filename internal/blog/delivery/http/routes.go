package http

import (
	"github.com/robertgouveia/portfolio/internal/blog"
	"net/http"
)

func MapRoutes(mux *http.ServeMux, h blog.Handler) {
	mux.HandleFunc("/blog", h.GetBlogs())
}
