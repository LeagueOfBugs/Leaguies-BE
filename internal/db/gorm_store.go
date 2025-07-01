package db

import "gorm.io/gorm"

func NewGormStore(db *gorm.DB) *Store {
	return &Store{
		Leagues: &GormLeagueStore{DB: db},
	}
}
