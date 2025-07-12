package models

import "time"

type Invitation struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"not null"`
	FromID    uint   `gorm:"not null"`
	ToID      uint   `gorm:"not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
