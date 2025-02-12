package server

import (
	bHttp "github.com/robertgouveia/portfolio/internal/blog/delivery/http"
	bUseCase "github.com/robertgouveia/portfolio/internal/blog/usecase"
	cHttp "github.com/robertgouveia/portfolio/internal/contact/delivery/http"
	pHttp "github.com/robertgouveia/portfolio/internal/project/delivery/http"
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

	buseCase := bUseCase.NewBlogUseCase()

	bHandler := bHttp.NewBlogHandler(s.errorLog, buseCase, s.Render)
	bHttp.MapRoutes(mux, bHandler)

	pHandler := pHttp.NewProjectHandler(s.Render)
	pHttp.MapRoutes(mux, pHandler)

	cHandler := cHttp.NewContactHandler(s.Render)
	cHttp.MapRoutes(mux, cHandler)

	return mux
}
