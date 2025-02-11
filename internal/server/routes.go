package server

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	images := http.FileServer(http.Dir("./public/static/img/"))
	mux.Handle("/img/", http.StripPrefix("/img/", images))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			s.NotFound(w)
			return
		}
		s.Render(w, nil, "base", "home", "nav")
	})

	return mux
}
