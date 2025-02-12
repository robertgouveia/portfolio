package http

import (
	"github.com/robertgouveia/portfolio/internal/blog"
	"log"
	"net/http"
)

type BlogHandler struct {
	uc     blog.UseCase
	logger *log.Logger
	render func(w http.ResponseWriter, data map[string]interface{}, layout string, page string, partial ...string)
}

func NewBlogHandler(log *log.Logger, uc blog.UseCase, r func(w http.ResponseWriter, data map[string]interface{}, layout string, page string, partial ...string)) blog.Handler {
	return &BlogHandler{
		uc:     uc,
		render: r,
		logger: log,
	}
}

func (h *BlogHandler) GetBlogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		blogs, err := h.uc.GetBlogs()
		if err != nil {
			h.logger.Println(err)
			h.render(w, nil, "base", "notfound", "nav") // TODO: Change to Error page
			return
		}

		h.render(w, map[string]interface{}{
			"blogs": blogs,
		}, "base", "blog", "nav")
	}
}
