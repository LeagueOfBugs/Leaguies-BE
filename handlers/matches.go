package handlers

import (
	"encoding/json"
	"fmt"
	"leaguies_backend/internal/db"
	"leaguies_backend/models"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type CreateMatchRequest struct {
	Name        string     `json:"name"`
	SeasonID    *uint      `json:"season_id"`
	HomeTeamID  *uint      `json:"home_team_id"`
	AwayTeamID  *uint      `json:"away_team_id"`
	HomeScore   *int       `json:"home_score"`
	AwayScore   *int       `json:"away_score"`
	Status      *string    `json:"status"`
	Location    *string    `json:"location"`
	ScheduledAt *time.Time `json:"scheduled_at"`
}
type UpdateMatchRequest struct {
	Name        *string    `json:"name"`
	SeasonID    *uint      `json:"season_id"`
	HomeTeamID  *uint      `json:"home_team_id"`
	AwayTeamID  *uint      `json:"away_team_id"`
	HomeScore   *int       `json:"home_score"`
	AwayScore   *int       `json:"away_score"`
	Status      *string    `json:"status"`
	Location    *string    `json:"location"`
	ScheduledAt *time.Time `json:"scheduled_at"`
}

func CreateMatch(w http.ResponseWriter, r *http.Request) {
	var req CreateMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	match := models.Match{
		Name:        req.Name,
		SeasonID:    req.SeasonID,
		HomeTeamID:  req.HomeTeamID,
		AwayTeamID:  req.AwayTeamID,
		HomeScore:   req.HomeScore,
		AwayScore:   req.AwayScore,
		Status:      req.Status,
		Location:    req.Location,
		ScheduledAt: req.ScheduledAt,
		
	}

	// print match
	fmt.Println(match)

	if err := db.DB.Create(&match).Error; err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create match", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

func GetMatch(w http.ResponseWriter, r *http.Request) {
	matchIdParam := chi.URLParam(r, "id")
	if matchIdParam == "" {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	var match models.Match
	if err := db.DB.First(&match, matchIdParam).Error; err != nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

func ListMatches(w http.ResponseWriter, r *http.Request) {
	var matches []models.Match
	if err := db.DB.Find(&matches).Error; err != nil {
		http.Error(w, "Failed to fetch matches", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(matches)
}

func DeleteMatch(w http.ResponseWriter, r *http.Request) {
	matchIdParam := chi.URLParam(r, "id")
	if matchIdParam == "" {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	var match models.Match
	if err := db.DB.First(&match, matchIdParam).Error; err != nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	if err := db.DB.Delete(&match).Error; err != nil {
		http.Error(w, "Failed to delete match", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)
}

func UpdateMatch(w http.ResponseWriter, r *http.Request) {
	// get id param
	matchIdParam := chi.URLParam(r, "id")
	if matchIdParam == "" {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	// decode request
	var req UpdateMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// check if match exists
	var match models.Match
	if err := db.DB.First(&match, matchIdParam).Error; err != nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	// update
	if req.Name != nil {
		match.Name = *req.Name
	}

	// save
	if err := db.DB.Save(&match).Error; err != nil {
		http.Error(w, "Failed to update match", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)
}
