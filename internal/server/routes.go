package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger/v2"
	
	_ "github.com/mohammad-gazali/job-board/docs"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Mount("/swagger", httpSwagger.WrapHandler)

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/", redirectToDocsHandler)
		r.Get("/hello", s.helloWorldHandler)
	})

	return router
}

func redirectToDocsHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/swagger/index.html", http.StatusSeeOther)
}