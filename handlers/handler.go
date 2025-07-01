package handlers

import "leaguies_backend/internal/db"

type Handler struct {
	League *LeagueHandler
}

func NewHandler(store *db.Store) *Handler {
	return &Handler{
		League : newLeagueHandler(store.Leagues),
	}
}
