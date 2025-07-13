package invite

import (
	"leaguies_backend/models"

	"gorm.io/gorm"
)

type InviteStoreInterface interface {
	Create(invite *models.Invite) error
	Delete(invite *models.Invite) error
	GetByID(id uint) (*models.Invite, error)
	List() ([]models.Invite, error)
}


type InviteStore struct {
	DB *gorm.DB
}

func (s *InviteStore) Create(invite *models.Invite) error {
	return s.DB.Create(invite).Error
}

func (s *InviteStore) Delete(invite *models.Invite) error {
	return s.DB.Delete(invite).Error
}

func (s *InviteStore) GetByID(id uint) (*models.Invite, error) {
	var invite models.Invite
	if err := s.DB.First(&invite, id).Error; err != nil {
		return nil, err
	}
	return &invite, nil
}

func (s *InviteStore) List() ([]models.Invite, error) {
	var invite []models.Invite
	if err := s.DB.Find(&invite).Error; err != nil {
		return nil, err
	}
	return invite, nil
}