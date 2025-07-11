package season

import (
	"leaguies_backend/models"

	"gorm.io/gorm"
)

type SeasonStoreInterface interface {
	GetAll() ([]models.Season, error)
	GetByID(id uint) (*models.Season, error)
	Create(season *models.Season) error
	Update(season *models.Season) error
	Delete(season *models.Season) error
}

type SeasonStore struct {
	DB *gorm.DB
}

func (s *SeasonStore) GetAll() ([]models.Season, error) {
	var seasons []models.Season
	if err := s.DB.Find(&seasons).Error; err != nil {
		return nil, err
	}
	return seasons, nil
}

func (s *SeasonStore) GetByID(id uint) (*models.Season, error) {
	var season models.Season
	if err := s.DB.First(&season, id).Error; err != nil {
		return nil, err
	}
	return &season, nil
}

func (s *SeasonStore) Create(season *models.Season) error {
	if err := s.DB.Create(season).Error; err != nil {
		return err
	}
	return nil
}

func (s *SeasonStore) Update(season *models.Season) error {
	if err := s.DB.Save(season).Error; err != nil {
		return err
	}
	return nil
}

func (s *SeasonStore) Delete(season *models.Season) error {
	if err := s.DB.Delete(season).Error; err != nil {
		return err
	}
	return nil
}