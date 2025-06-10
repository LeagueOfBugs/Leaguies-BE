package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"leaguies_backend/internal/db"
	"leaguies_backend/models"
	"github.com/go-chi/chi/v5"
)

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
