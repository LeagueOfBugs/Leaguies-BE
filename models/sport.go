package models

type Sport struct {
	ID uint     `gorm:"primaryKey"`
	Name    string   `gorm:"uniqueIndex;not null"`
	// fix issue where league gets created but not added to sport
	Leagues []League `gorm:"foreignKey:SportID"`
}
