package contact

import "net/http"

type Handler interface {
	GetContact() http.HandlerFunc
}
