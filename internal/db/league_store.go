package db

import "leaguies_backend/models"

type LeagueStore interface {
	Create(league *models.League) error
	Update(league *models.League) error
	Delete(league *models.League) error
	GetByID(id uint) (*models.League, error)
	List() ([]models.League, error)
}
