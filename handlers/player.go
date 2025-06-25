package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"leaguies_backend/internal/db"
	"leaguies_backend/models"
	"github.com/go-chi/chi/v5"
)

type CreatePlayerRequest struct {
	Gender string `json:"gender"`
	Position string `json:"position"`
	UserID uint `json:"user_id"`
	SportID uint `json:"sport_id"`
}

type UpdatePlayerRequest struct {
	Gender *string `json:"gender"`
	Position *string `json:"position"`
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	// Get player ID from URL params, e.g., /players/{id}
	playerIDStr := chi.URLParam(r, "id")
	playerID, err := strconv.ParseUint(playerIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	var player models.Player
	// Preload User to include related user info if you want
	if err := db.DB.Preload("User").First(&player, playerID).Error; err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

func ListPlayers(w http.ResponseWriter, r *http.Request) {
	var players []models.Player
	if err := db.DB.Find(&players).Error; err != nil {
		http.Error(w, "Failed to fetch players", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(players)
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var req CreatePlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	player := models.Player{
		Gender: req.Gender,
		Position: req.Position,
		UserID: req.UserID,
		SportID: req.SportID,
	}

	if err := db.DB.Create(&player).Error; err != nil {
		http.Error(w, "Failed to create player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	// get id param
	playerIdParam := chi.URLParam(r, "id")
	if playerIdParam == "" {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	// decode request
	var req UpdatePlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// check if player exists
	var player models.Player
	if err := db.DB.First(&player, playerIdParam).Error; err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	// update
	if req.Position != nil {
		player.Position = *req.Position
	}

	// save
	if err := db.DB.Save(&player).Error; err != nil {
		http.Error(w, "Failed to update player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	// get id param
	playerIdParam := chi.URLParam(r, "id")
	if playerIdParam == "" {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	var player models.Player
	if err := db.DB.First(&player, playerIdParam).Error; err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	if err := db.DB.Delete(&player).Error; err != nil {
		http.Error(w, "Failed to delete player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}


