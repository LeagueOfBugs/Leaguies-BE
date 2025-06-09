package router

import (
	"leaguies_backend/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health-check", handlers.HealthCheck)
	r.Post("/register", handlers.Register)
	
	return r
}
