package blog

import "net/http"

type Handler interface {
	GetBlogs() http.HandlerFunc
	GetBlog() http.HandlerFunc
}
