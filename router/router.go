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
			r.Get("/", h.League.List)
			r.Get("/{id}", h.League.GetByID)
			r.Post("/create", h.League.Create)
			r.Put("/{id}/update", h.League.Update)
			r.Delete("/{id}/delete", h.League.Delete)
		})

		// season routes
		r.Route("/season", func(r chi.Router) {
			r.Get("/", h.Season.List)
			r.Get("/{id}", h.Season.GetByID)
			r.Post("/create", h.Season.Create)
			r.Put("/{id}", h.Season.Update)
			r.Delete("/{id}", h.Season.Delete)
		})

		// team routes
		r.Route("/team", func(r chi.Router) {
			r.Post("/create", h.Team.Create)
			r.Get("/", h.Team.List)
			r.Get("/{id}", h.Team.GetByID)
			r.Put("/{id}", h.Team.Update)
			r.Delete("/{id}", h.Team.Delete)
		})

		r.Route("/match", func(r chi.Router) {
			r.Post("/create", h.Match.Create)
			r.Get("/", h.Match.List)
			r.Get("/{id}", h.Match.GetByID)
			r.Put("/{id}", h.Match.Update)
			r.Delete("/{id}", h.Match.Delete)
		})

		r.Route("/invite", func(r chi.Router) {
			r.Post("/create", h.Invite.Create)
			r.Get("/", h.Invite.List)
			r.Get("/{id}", h.Invite.GetByID)
			r.Delete("/{id}", h.Invite.Delete)
		})
	})

	return r
}
