package router

import (
	"leaguies_backend/handlers"
	"leaguies_backend/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health-check", handlers.HealthCheck)
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)
	
	// protected routes
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Get("/user", handlers.Me)
		r.Get("/player/{id}", handlers.GetPlayer)
	})

	return r
}
