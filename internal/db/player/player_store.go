package player

import (
	"leaguies_backend/models"

	"gorm.io/gorm"
)

type PlayerStoreInterface interface {
	Create(player *models.Player) error
	Update(player *models.Player) error
	Delete(player *models.Player) error
	GetByID(id uint) (*models.Player, error)
	List() ([]models.Player, error)
}

type PlayerStore struct {
	DB *gorm.DB
}

func (s *PlayerStore) Create(player *models.Player) error {
	return s.DB.Create(player).Error
}

func (s *PlayerStore) Update(player *models.Player) error {
	return s.DB.Save(player).Error
}

func (s *PlayerStore) Delete(player *models.Player) error {
	return s.DB.Delete(player).Error
}

func (s *PlayerStore) GetByID(id uint) (*models.Player, error) {
	var player models.Player
	if err := s.DB.First(&player, id).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func (s *PlayerStore) List() ([]models.Player, error) {
	var players []models.Player
	if err := s.DB.Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}
