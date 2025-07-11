package handlers

import (
	"leaguies_backend/handlers/league"
	"leaguies_backend/handlers/season"
	"leaguies_backend/internal/db"
)

type Handler struct {
	League *league.LeagueHandler
	Season *season.SeasonHandler
}

func NewHandler(store *db.Store) *Handler {
	return &Handler{
		League: league.NewLeagueHandler(store.Leagues),
		Season: season.NewSeasonHandler(store.Seasons),
	}
}
