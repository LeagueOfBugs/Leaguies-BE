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

		// user routes
		r.Route("/user", func(r chi.Router) {
			r.Get("/", handlers.Me)
		})

		// players routes
		r.Route("/player", func(r chi.Router) {
			r.Get("/{id}", handlers.GetPlayer)
		})

		// leagues routes
		r.Route("/league", func(r chi.Router) {
			r.Get("/", handlers.ListLeagues)
			r.Get("/{id}", handlers.GetLeague)
			r.Post("/create", handlers.CreateLeague)
			r.Put("/{id}/update", handlers.UpdateLeague)
			r.Delete("/{id}/delete", handlers.DeleteLeague)
		})

		// season routes
		// r.Route("/season", func(r chi.Router) {
		// 	r.Post("/create", handlers.CreateSeason)
		// })
	})

	return r
}
