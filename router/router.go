package router

import (
	"leaguies_backend/handlers"
	"leaguies_backend/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
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
			r.Get("/", handlers.ListPlayers)
			r.Post("/create", handlers.CreatePlayer)
			r.Put("/{id}/update", handlers.UpdatePlayer)
			r.Delete("/{id}/delete", handlers.DeletePlayer)
		})

		// leagues routes
		r.Route("/league", func(r chi.Router) {
			r.Get("/", handlers.ListLeagues)
			r.Get("/{id}", handlers.GetLeague)
			r.Post("/create", h.League.Create)
			r.Put("/{id}/update", handlers.UpdateLeague)
			r.Delete("/{id}/delete", handlers.DeleteLeague)
		})

		// season routes
		r.Route("/season", func(r chi.Router) {
			r.Post("/create", handlers.CreateSeason)
			r.Get("/", handlers.ListSeason)
			r.Get("/{id}", handlers.GetSeason)
			r.Put("/{id}", handlers.UpdateSeason)
			r.Delete("/{id}", handlers.DeleteSeason)
		})

		// team routes
		r.Route("/team", func(r chi.Router) {
			r.Post("/create", handlers.CreateTeam)
			r.Get("/", handlers.ListTeams)
			r.Get("/{id}", handlers.GetTeam)
			r.Put("/{id}", handlers.UpdateTeam)
			r.Delete("/{id}", handlers.DeleteTeam)
		})

		r.Route("/match", func(r chi.Router) {
			r.Post("/create", handlers.CreateMatch)
			r.Get("/", handlers.ListMatches)
			r.Get("/{id}", handlers.GetMatch)
			r.Put("/{id}", handlers.UpdateMatch)
			r.Delete("/{id}", handlers.DeleteMatch)
		})
	})

	return r
}
