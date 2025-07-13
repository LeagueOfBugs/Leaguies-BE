package models

import "time"

type Invite struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"not null"`
	FromID    uint   `gorm:"not null"`
	ToID      uint   `gorm:"not null"`
	Status    string `gorm:"type:varchar(20)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
