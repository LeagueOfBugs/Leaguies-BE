package models

import "time"

type Match struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	SeasonID *uint   `gorm:"not null"`
	Season   Season `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	HomeTeamID *uint `gorm:"not null"`
	HomeTeam   Team `gorm:"foreignKey:HomeTeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	AwayTeamID  *uint `gorm:"not null"`
	AwayTeam    Team `gorm:"foreignKey:AwayTeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HomeScore   *int
	AwayScore   *int
	Status      *string `gorm:"type:varchar(20)"`
	Location    *string
	ScheduledAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
