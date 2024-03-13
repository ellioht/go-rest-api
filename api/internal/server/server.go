package server

import (
	"context"
	"fmt"
	"github.com/ellioht/go-rest-api/config"
	"github.com/ellioht/go-rest-api/pkg/atomics"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	Config       *config.Config
	ShuttingDown *atomics.AtomicBool
	Router       *chi.Mux
	Deps         *dependencies
}

func (s *Server) Init() error {
	return nil
}

func (s *Server) Run(ct context.Context) error {
	ctx, cancel := context.WithCancel(ct)
	defer cancel()

	if err := s.Setup(); err != nil {
		return err
	}

	port := fmt.Sprintf(":%v", s.Config.Port)
	server := http.Server{
		Addr:    port,
		Handler: s.Router,
	}

	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			return
		}
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Setup() error {
	if err := s.SetupDeps(); err != nil {
		return err
	}

	if err := s.SetupHandlers(); err != nil {
		return err
	}

	return nil
}
