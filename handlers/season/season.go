package season

import (
	"encoding/json"
	"leaguies_backend/internal/db"
	"leaguies_backend/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// type CreateSeasonRequest struct {
// 	Name      string    `json:"name"`
// 	LeagueID  uint      `json:"league_id"`
// 	StartDate time.Time `json:"start_date"`
// 	EndDate   time.Time `json:"end_date"`
// }

// type UpdateSeasonRequest struct {
// 	Name      *string    `json:"name"`
// 	LeagueID  *uint      `json:"league_id"`
// 	StartDate *time.Time `json:"start_date"`
// 	EndDate   *time.Time `json:"end_date"`
// }

func CreateSeason(w http.ResponseWriter, r *http.Request) {
	var req CreateSeasonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	season := models.Season{
		Name:      req.Name,
		LeagueID:  req.LeagueID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	if err := db.DB.Create(&season).Error; err != nil {
		http.Error(w, "Failed to create season", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(season)
}

func UpdateSeason(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var season models.Season
	if err := db.DB.First(&season, id).Error; err != nil {
		http.Error(w, "Season not found", http.StatusNotFound)
		return
	}

	var req UpdateSeasonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name != nil {
		season.Name = *req.Name
	}

	if req.LeagueID != nil {
		season.LeagueID = *req.LeagueID
	}

	if req.StartDate != nil {
		season.StartDate = *req.StartDate
	}

	if req.EndDate != nil {
		season.EndDate = *req.EndDate
	}

	if err := db.DB.Save(&season).Error; err != nil {
		http.Error(w, "Failed to update season", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(season)
}

func DeleteSeason(w http.ResponseWriter, r *http.Request) {
	seasonIdParam := chi.URLParam(r, "id")
	if seasonIdParam == "" {
		http.Error(w, "Invalid season ID", http.StatusBadRequest)
		return
	}


	if err := db.DB.First(&models.Season{}, seasonIdParam).Error; err != nil {
		http.Error(w, "Season not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetSeason(w http.ResponseWriter, r *http.Request) {
	seasonIdParam := chi.URLParam(r, "id")
	if seasonIdParam == "" {
		http.Error(w, "Invalid season ID", http.StatusBadRequest)
		return
	}

	var season models.Season
	if err := db.DB.First(&season, seasonIdParam).Error; err != nil {
		http.Error(w, "Season not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(season)
}

func ListSeason(w http.ResponseWriter, r *http.Request) {
	var seasons []models.Season
	if err := db.DB.Find(&seasons).Error; err != nil {
		http.Error(w, "Failed to fetch seasons", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(seasons)
}
