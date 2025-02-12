package http

import (
	"github.com/robertgouveia/portfolio/internal/blog"
	"net/http"
)

func MapRoutes(mux *http.ServeMux, h blog.Handler) {
	mux.HandleFunc("/blogs", h.GetBlogs())
	mux.HandleFunc("/blog/", h.GetBlog())
}
