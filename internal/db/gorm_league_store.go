package db

import (
	"leaguies_backend/models"

	"gorm.io/gorm"
)

type GormLeagueStore struct {
	DB *gorm.DB
}

func (s *GormLeagueStore) Create(league *models.League) error {
	return s.DB.Create(league).Error
}

func (s *GormLeagueStore) Update(league *models.League) error {
	return s.DB.Save(league).Error
}

func (s *GormLeagueStore) Delete(league *models.League) error {
	return s.DB.Delete(league).Error
}

func (s *GormLeagueStore) GetByID(id uint) (*models.League, error) {
	var league models.League
	if err := s.DB.First(&league, id).Error; err != nil {
		return nil, err
	}
	return &league, nil
}

func (s *GormLeagueStore) List() ([]models.League, error) {
	var leagues []models.League
	if err := s.DB.Find(&leagues).Error; err != nil {
		return nil, err
	}
	return leagues, nil
}
