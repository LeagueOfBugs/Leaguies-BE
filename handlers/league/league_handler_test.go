package league

import (
	"bytes"
	"encoding/json"
	// "errors"

	// "errors"
	"leaguies_backend/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockLeagueStore is a lightweight mock implementation of LeagueStore
type mockLeagueStore struct {
	CreateFn func(league *models.League) error
	UpdateFn func(league *models.League) error
	DeleteFn func(league *models.League) error
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

	handler := NewLeagueHandler(mockStore)

	tests := []struct {
		name           string
		input          map[string]interface{}
		setupMock      func()
		expectedStatus int
		expectedName   string
		expectedError  string
	}{
		{
			name: "Successful creation",
			input: map[string]interface{}{
				"name":     "Test League",
				"sport_id": 2,
			},
			setupMock: func() {
				mockStore.CreateFn = func(league *models.League) error {
					league.ID = 1
					return nil
				}
			},
			expectedStatus: http.StatusCreated,
			expectedName:   "Test League",
		},
		{
			name: "Missing name field",
			input: map[string]interface{}{
				"sport_id": 2,
			},
			setupMock: func() {
				mockStore.CreateFn = func(league *models.League) error {
					t.Fatal("Create should not be called for invalid input")
					return nil
				}
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Missing required fields: name and sportId",
		},
		// {
		// 	name: "DB error during creation",
		// 	input: map[string]interface{}{
		// 		"name":     "Test League",
		// 		"sport_id": 2,
		// 	},
		// 	setupMock: func() {
		// 		mockStore.CreateFn = func(league *models.League) error {
		// 			return errors.New("database error")
		// 		}
		// 	},
		// 	expectedStatus: http.StatusInternalServerError,
		// 	expectedError:  "Failed to create league",
		// },
		// {
		// 	name:  "Invalid JSON input",
		// 	input: nil,
		// 	setupMock: func() {
		// 		mockStore.CreateFn = func(league *models.League) error { return nil }
		// 	},
		// 	expectedStatus: http.StatusBadRequest,
		// 	expectedError:  "invalid request payload",
		// },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			var req *http.Request

			if tc.name == "Invalid JSON input" {
				req = httptest.NewRequest(http.MethodPost, "/leagues", bytes.NewReader([]byte("{invalid json")))
			} else {
				bodyBytes, err := json.Marshal(tc.input)
				if err != nil {
					t.Fatalf("Failed to marshal input: %v", err)
				}
				req = httptest.NewRequest(http.MethodPost, "/leagues", bytes.NewReader(bodyBytes))
			}

			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			handler.Create(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectedStatus, res.StatusCode)

			if res.StatusCode == http.StatusCreated {
				var created models.League
				err := json.NewDecoder(res.Body).Decode(&created)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedName, created.Name)
			} else {
				var errResp struct {
					Error string `json:"error"`
				}
				err := json.NewDecoder(res.Body).Decode(&errResp)
				assert.NoError(t, err)
				assert.Contains(t, errResp.Error, tc.expectedError)
			}
		})
	}
}
