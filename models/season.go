package models

import "time"

type Season struct {
	ID uint `gorm:"primaryKey"`
	// ensures that season name acan exist in multiple leagues, bit not more than once per league
	Name      string `gorm:"not null;uniqueIndex:idx_league_season"`
	LeagueID  uint   `gorm:"not null;uniqueIndex:idx_league_season"`
	League    League `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
