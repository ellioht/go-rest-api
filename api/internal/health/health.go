package health

import (
	"fmt"
	"github.com/ellioht/go-rest-api/pkg/atomics"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	shuttingDown *atomics.AtomicBool
}

func NewHandler(r chi.Router, shuttingDown *atomics.AtomicBool) *Handler {
	h := &Handler{
		shuttingDown: shuttingDown,
	}
	h.SetupRoutes(r)
	return h
}

func (h *Handler) SetupRoutes(router chi.Router) {
	fmt.Println("setting up routes for health...")
	router.Group(func(r chi.Router) {
		r.Get("/health", h.GetHealth)
	})
}

func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	if h.shuttingDown.Get() {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	_, err := w.Write([]byte("OK"))
	if err != nil {
		fmt.Println(err)
	}
}
