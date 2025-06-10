package handlers

import (
	"encoding/json"
	"leaguies_backend/internal/db"
	"leaguies_backend/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CreateLeagueRequest struct {
	Name    string `json:"name"`
	SportId uint   `json:"sport_id"`
}

func CreateLeague(w http.ResponseWriter, r *http.Request) {
	var req CreateLeagueRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	league := models.League{
		Name:    req.Name,
		SportID: req.SportId,
	}

	if err := db.DB.Create(&league).Error; err != nil {
		http.Error(w, "Failed to create league", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(league)
}

type UpdateLeagueRequest struct {
	Name    *string `json:"name"`
	SportId *uint   `json:"sport_id"`
}

func UpdateLeague(w http.ResponseWriter, r *http.Request) {
	// get id param
	leagueIdParam := chi.URLParam(r, "id")
	if leagueIdParam == "" {
		http.Error(w, "Invalid league ID", http.StatusBadRequest)
		return
	}

	// decode request
	var req UpdateLeagueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// check if league exists
	var league models.League
	if err := db.DB.First(&league, leagueIdParam).Error; err != nil {
		http.Error(w, "League not found", http.StatusNotFound)
		return
	}

	// update
	if req.Name != nil {
		league.Name = *req.Name
	}

	if req.SportId != nil {
		league.SportID = *req.SportId
	}

	// save
	if err := db.DB.Save(&league).Error; err != nil {
		http.Error(w, "Failed to update league", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(league)
}

func DeleteLeague(w http.ResponseWriter, r *http.Request) {
	leagueIdParam := chi.URLParam(r, "id")
	if leagueIdParam == "" {
		http.Error(w, "Invalid league ID", http.StatusBadRequest)
		return
	}

	var league models.League
	if err := db.DB.First(&league, leagueIdParam).Error; err != nil {
		http.Error(w, "League not found", http.StatusNotFound)
		return
	}

	if err := db.DB.Delete(&league).Error; err != nil {
		http.Error(w, "Failed to delete league", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(league)
}

func GetLeague(w http.ResponseWriter, r *http.Request) {
	leagueIdParam := chi.URLParam(r, "id")
	if leagueIdParam == "" {
		http.Error(w, "Invalid league ID", http.StatusBadRequest)
		return
	}

	var league models.League
	if err := db.DB.First(&league, leagueIdParam).Error; err != nil {
		http.Error(w, "League not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(league)
}

func ListLeagues(w http.ResponseWriter, r *http.Request) {
	var leagues []models.League
	if err := db.DB.Find(&leagues).Error; err != nil {
		http.Error(w, "Failed to fetch leagues", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(leagues)
}
