package project

import "net/http"

type Handler interface {
	GetProjects() http.HandlerFunc
}
