package handlers

import (
	"encoding/json"
	"leaguies_backend/internal/db"
	"leaguies_backend/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CreateTeamRequest struct {
	Name     string `json:"name"`
	LeagueId *uint  `json:"league_id"`
}

type UpdateTeamRequest struct {
	Name     *string `json:"name"`
	LeagueId *uint   `json:"league_id"`
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var req CreateTeamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	team := models.Team{
		Name:     req.Name,
		LeagueId: req.LeagueId,
	}

	if err := db.DB.Create(&team).Error; err != nil {
		http.Error(w, "Failed to create team", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var team models.Team
	if err := db.DB.First(&team, id).Error; err != nil {
		http.Error(w, "Team not found", http.StatusNotFound)
		return
	}

	var req UpdateTeamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name != nil {
		team.Name = *req.Name
	}

	team.LeagueId = req.LeagueId

	if err := db.DB.Save(&team).Error; err != nil {
		http.Error(w, "Failed to update team", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := db.DB.Delete(&models.Team{}, id).Error; err != nil {
		http.Error(w, "Failed to delete team", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var team models.Team
	if err := db.DB.First(&team, id).Error; err != nil {
		http.Error(w, "Team not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}

func ListTeams(w http.ResponseWriter, r *http.Request) {
	var teams []models.Team
	if err := db.DB.Find(&teams).Error; err != nil {
		http.Error(w, "Failed to fetch teams", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teams)
}
