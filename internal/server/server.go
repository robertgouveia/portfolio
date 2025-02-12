package server

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	addr          string
	infoLog       *log.Logger
	errorLog      *log.Logger
	errorChan     chan error
	shutdownChan  chan struct{}
	templateCache map[string]*template.Template
	env           string
}

func NewServer(addr string, env string) *Server {
	return &Server{
		addr:          addr,
		infoLog:       log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog:      log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		errorChan:     make(chan error, 1),
		shutdownChan:  make(chan struct{}),
		templateCache: make(map[string]*template.Template),
		env:           env,
	}
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.routes(),
	}

	s.infoLog.Printf("listening on %s", s.addr)

	if s.env == "production" {
		go func() {
			fmt.Println("Redirecting HTTP to HTTPS on port 80...")
			err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "https://websitech.uk", http.StatusMovedPermanently)
			}))
			if err != nil {
				fmt.Println("Error starting HTTP server:", err)
			}
		}()

		go func() {
			if err := srv.ListenAndServeTLS("/etc/letsencrypt/live/websitech.uk/fullchain.pem", "/etc/letsencrypt/live/websitech.uk/privkey.pem"); err != nil && err != http.ErrServerClosed {
				s.errorLog.Println(err)
				s.errorChan <- err
			}
		}()
	} else {
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				s.errorLog.Println(err)
				s.errorChan <- err
			}
		}()
	}

	select {
	case err := <-s.errorChan:
		return err
	case <-s.shutdownChan:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			s.errorLog.Println(err)
			return err
		}

		return nil
	}
}

func (s *Server) Shutdown() {
	close(s.shutdownChan)
}

func (s *Server) ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	s.infoLog.Println("Shutting down server...")
	s.Shutdown()
	s.infoLog.Println("Server gracefully stopped")
}

func (s *Server) Render(w http.ResponseWriter, data *interface{}, layout string, page string, partial ...string) {
	if s.templateCache[page] != nil && s.env == "production" {
		s.infoLog.Printf("Cache hit: %s", page)
		s.templateCache[page].Execute(w, data)
		return
	}

	s.infoLog.Printf("Initial paint: %s", page)
	var partials []string
	partials = append(partials, fmt.Sprintf("public/templates/%s.tmpl", layout))
	partials = append(partials, fmt.Sprintf("public/templates/%s.tmpl", page))
	for _, filename := range partial {
		partials = append(partials, fmt.Sprintf("public/templates/partials/%s.tmpl", filename))
	}

	tmpl, err := template.ParseFiles(partials...)
	if err != nil {
		s.errorLog.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.templateCache[page] = tmpl

	tmpl.Execute(w, data)
}

func (s *Server) NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	s.Render(w, nil, "base", "notfound", "nav")
}
