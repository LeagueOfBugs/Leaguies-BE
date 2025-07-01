package handlers

import (
	"bytes"
	"encoding/json"
	"leaguies_backend/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

// mockLeagueStore is a lightweight mock implementation of LeagueStore
type mockLeagueStore struct {
	CreateFn func(league *models.League) error
}

func (m *mockLeagueStore) Create(league *models.League) error {
	return m.CreateFn(league)
}

func (m *mockLeagueStore) Update(*models.League) error { return nil }
func (m *mockLeagueStore) Delete(*models.League) error { return nil }
func (m *mockLeagueStore) GetByID(id uint) (*models.League, error) {
	return &models.League{ID: id, Name: "Mock"}, nil
}
func (m *mockLeagueStore) List() ([]models.League, error) { return []models.League{}, nil }

func TestCreateLeague(t *testing.T) {
	mockStore := &mockLeagueStore{
		CreateFn: func(league *models.League) error {
			league.ID = 1 
			return nil
		},
	}

	handler := newLeagueHandler(mockStore)

	body := map[string]any{
		"name":     "Test League",
		"sport_id": 2,
	}

	b, err := json.Marshal(body)

	if err != nil {
		log.Fatal("Failed to marshal:", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/leagues", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	// create response recorder
	w := httptest.NewRecorder()

	handler.Create(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var created models.League
	json.NewDecoder(res.Body).Decode(&created)
	assert.Equal(t, "Test League", created.Name)
	assert.Equal(t, uint(2), created.SportID)
}
