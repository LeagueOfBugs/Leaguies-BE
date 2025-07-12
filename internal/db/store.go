package db

import "gorm.io/gorm"
import "leaguies_backend/internal/db/league"
import "leaguies_backend/internal/db/season"
import "leaguies_backend/internal/db/team"

type Store struct {
	Leagues league.LeagueStoreInterface
	Seasons season.SeasonStoreInterface
	Teams   team.TeamStoreInterface
	// Add: Players PlayerStore, Matches MatchStore, etc.
}


func NewStore(db *gorm.DB) *Store {
	return &Store{
		Leagues: &league.LeagueStore{DB: db},
		Seasons: &season.SeasonStore{DB: db},
		Teams:   &team.TeamStore{DB: db},
	}
}
