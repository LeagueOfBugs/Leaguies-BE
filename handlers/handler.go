package handlers

import (
	"leaguies_backend/handlers/league"
	"leaguies_backend/handlers/season"
	"leaguies_backend/handlers/team"
	"leaguies_backend/internal/db"
)

type Handler struct {
	League *league.LeagueHandler
	Season *season.SeasonHandler
	Team  *team.TeamHandler
}

func NewHandler(store *db.Store) *Handler {
	return &Handler{
		League: league.NewLeagueHandler(store.Leagues),
		Season: season.NewSeasonHandler(store.Seasons),
		Team:  team.NewTeamHandler(store.Teams),
	}
}
