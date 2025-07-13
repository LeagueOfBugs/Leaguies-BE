package handlers

import (
	"leaguies_backend/handlers/invite"
	"leaguies_backend/handlers/league"
	"leaguies_backend/handlers/match"
	"leaguies_backend/handlers/season"
	"leaguies_backend/handlers/team"
	"leaguies_backend/internal/db"
)

type Handler struct {
	League *league.LeagueHandler
	Season *season.SeasonHandler
	Team  *team.TeamHandler
	Match *match.MatchHandler
	Invite *invite.InviteHandler
}

func NewHandler(store *db.Store) *Handler {
	return &Handler{
		League: league.NewLeagueHandler(store.Leagues),
		Season: season.NewSeasonHandler(store.Seasons),
		Team:  team.NewTeamHandler(store.Teams),
		Match: match.NewMatchHandler(store.Matches, store.Teams),
		Invite: invite.NewInviteHandler(store.Invites),
	}
}
