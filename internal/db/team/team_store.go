package team

import (
	"leaguies_backend/models"

	"gorm.io/gorm"
)

type TeamStoreInterface interface {
	Create(team *models.Team) error
	Update(team *models.Team) error
	Delete(team *models.Team) error
	GetByID(id uint) (*models.Team, error)
	List() ([]models.Team, error)
}

type TeamStore struct {
	DB *gorm.DB
}

func (s *TeamStore) Create(team *models.Team) error {
	return s.DB.Create(team).Error
}

func (s *TeamStore) Update(team *models.Team) error {
	return s.DB.Save(team).Error
}

func (s *TeamStore) Delete(team *models.Team) error {
	return s.DB.Delete(team).Error
}

func (s *TeamStore) GetByID(id uint) (*models.Team, error) {
	var team models.Team
	if err := s.DB.First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (s *TeamStore) List() ([]models.Team, error) {
	var teams []models.Team
	if err := s.DB.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
