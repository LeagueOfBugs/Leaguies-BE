package mocks

import "leaguies_backend/models"

type mockLeagueStore struct {
	CreateFn func(league *models.League) error
}

func (m *mockLeagueStore) Create(league *models.League) error {
	return m.CreateFn(league)
}

func (m *mockLeagueStore) Update(*models.League) error { return nil }
func (m *mockLeagueStore) Delete(*models.League) error { return nil }
func (m *mockLeagueStore) GetByID(id uint) (*models.League, error) {
	return &models.League{ID: id, Name: "Mock League"}, nil
}
func (m *mockLeagueStore) List() ([]models.League, error) { return []models.League{}, nil }
