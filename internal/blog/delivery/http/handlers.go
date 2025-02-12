package http

import (
	"github.com/robertgouveia/portfolio/internal/blog"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		search := r.URL.Query().Get("search")

		blogs, err := h.uc.GetBlogs(search)
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

func (h *BlogHandler) GetBlog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		param := strings.TrimPrefix(path, "/blog/")
		id, err := strconv.Atoi(param)
		if err != nil || param == " " {
			h.render(w, nil, "base", "notfound", "nav") // TODO: Change to Error page
			return
		}

		b, err := h.uc.GetBlog(id)
		if err != nil {
			h.logger.Println(err)
			h.render(w, nil, "base", "notfound", "nav") // TODO: Change to Error page
			return
		}

		h.render(w, map[string]interface{}{
			"blog": b,
		}, "base", "blogPage", "nav")
	}
}
