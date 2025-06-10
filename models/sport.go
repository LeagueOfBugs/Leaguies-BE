package models

type Sport struct {
	ID uint     `gorm:"primaryKey"`
	Name    string   `gorm:"uniqueIndex;not null"`
	Leagues []League `gorm:"foreignKey:SportID"`
}
