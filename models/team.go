package models

import "time"

type Team struct {
	ID    uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex;not null"`
	LeagueId  *uint
	League    *League `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
