package match

import (
	"leaguies_backend/models"

	"gorm.io/gorm"
)

type MatchStoreInterface interface {
	List() ([]models.Match, error)
	GetByID(id uint) (*models.Match, error)
	Create(match *models.Match) error
	Update(match *models.Match) error
	Delete(match *models.Match) error
}

type MatchStore struct {
	DB *gorm.DB
}

func (s *MatchStore) List() ([]models.Match, error) {
	var matches []models.Match
	if err := s.DB.Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

func (s *MatchStore) GetByID(id uint) (*models.Match, error) {
	var matches models.Match
	if err := s.DB.First(&matches, id).Error; err != nil {
		return nil, err
	}
	return &matches, nil
}

func (s *MatchStore) Create(matches *models.Match) error {
	if err := s.DB.Create(matches).Error; err != nil {
		return err
	}
	return nil
}

func (s *MatchStore) Update(matches *models.Match) error {
	if err := s.DB.Save(matches).Error; err != nil {
		return err
	}
	return nil
}

func (s *MatchStore) Delete(matches *models.Match) error {
	if err := s.DB.Delete(matches).Error; err != nil {
		return err
	}
	return nil
}