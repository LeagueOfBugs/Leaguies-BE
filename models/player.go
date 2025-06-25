package models

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	ID  uint   `gorm:"primaryKey"`
	Gender    string `gorm:"type:varchar(10)"`
	Position  string `gorm:"type:varchar(50)"`
	UserID    uint   `gorm:"not null;unique"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SportID   uint   `gorm:"not null"`
	Sport     Sport  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
