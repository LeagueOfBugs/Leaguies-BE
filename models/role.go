package models

import (
	"time"
)

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null" json:"name"`
	Users     []User    `gorm:"many2many:user_roles"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
