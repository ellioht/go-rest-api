package server

import (
	"github.com/ellioht/go-rest-api/internal/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type dependencies struct {
}

func (s *Server) SetupDeps() error {
	var deps dependencies
	s.Deps = &deps
	return nil
}

func (s *Server) SetupHandlers() error {

	s.Router = chi.NewRouter()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	s.Router.Use(crs.Handler)

	s.Router.Route("/api/v1", func(r chi.Router) {
		health.NewHandler(r, s.ShuttingDown)
	})
	return nil
}
