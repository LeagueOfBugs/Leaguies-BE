package league

import (
	"leaguies_backend/models"

	"gorm.io/gorm"
)

type LeagueStoreInterface interface {
	Create(league *models.League) error
	Update(league *models.League) error
	Delete(league *models.League) error
	GetByID(id uint) (*models.League, error)
	List() ([]models.League, error)
}


type LeagueStore struct {
	DB *gorm.DB
}

func (s *LeagueStore) Create(league *models.League) error {
	return s.DB.Create(league).Error
}

func (s *LeagueStore) Update(league *models.League) error {
	return s.DB.Save(league).Error
}

func (s *LeagueStore) Delete(league *models.League) error {
	return s.DB.Delete(league).Error
}

func (s *LeagueStore) GetByID(id uint) (*models.League, error) {
	var league models.League
	if err := s.DB.First(&league, id).Error; err != nil {
		return nil, err
	}
	return &league, nil
}

func (s *LeagueStore) List() ([]models.League, error) {
	var leagues []models.League
	if err := s.DB.Find(&leagues).Error; err != nil {
		return nil, err
	}
	return leagues, nil
}
