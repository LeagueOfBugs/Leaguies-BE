package models

import "time"

type League struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex;not null"`
	SportID   uint   `gorm:"not null"`
	Sport     Sport  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
