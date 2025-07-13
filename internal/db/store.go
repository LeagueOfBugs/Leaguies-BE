package db

import (
	"leaguies_backend/internal/db/league"
	"leaguies_backend/internal/db/match"
	"leaguies_backend/internal/db/season"
	"leaguies_backend/internal/db/team"
	"leaguies_backend/internal/db/invite"

	"gorm.io/gorm"
)

type Store struct {
	Leagues league.LeagueStoreInterface
	Seasons season.SeasonStoreInterface
	Teams   team.TeamStoreInterface
	Matches match.MatchStoreInterface
	Invites invite.InviteStoreInterface
	// Add: Players PlayerStore, Matches MatchStore, etc.
}


func NewStore(db *gorm.DB) *Store {
	return &Store{
		Leagues: &league.LeagueStore{DB: db},
		Seasons: &season.SeasonStore{DB: db},
		Teams:   &team.TeamStore{DB: db},
		Matches: &match.MatchStore{DB: db},
		Invites: &invite.InviteStore{DB: db},
	}
}
