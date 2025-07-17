package player

import (
	"encoding/json"
	"net/http"

	"leaguies_backend/internal/db"
	"leaguies_backend/internal/db/player"
	"leaguies_backend/models"

	"github.com/go-chi/chi/v5"
)

type PlayerHandler struct {
	store player.PlayerStoreInterface
}

type CreatePlayerRequest struct {
	Gender   string `json:"gender"`
	Position string `json:"position"`
	UserID   uint   `json:"user_id"`
	SportID  uint   `json:"sport_id"`
}

type UpdatePlayerRequest struct {
	Gender   *string `json:"gender"`
	Position *string `json:"position"`
}

func NewPlayerHandler(store player.PlayerStoreInterface) *PlayerHandler {
	return &PlayerHandler{
		store: store,
	}
}

func (s *PlayerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreatePlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	player := models.Player{
		Gender:   req.Gender,
		Position: req.Position,
		UserID:   req.UserID,
		SportID:  req.SportID,
	}

	if err := s.store.Create(&player).Error; err != nil {
		http.Error(w, "Failed to create player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

func (s *PlayerHandler) List(w http.ResponseWriter, r *http.Request) {
	var players []models.Player
	if err := db.DB.Find(&players).Error; err != nil {
		http.Error(w, "Failed to fetch players", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(players)
}

func (s *PlayerHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	if err := s.store.Update(&player); err != nil {
		http.Error(w, "Failed to update player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

func (s *PlayerHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	if err := s.store.Delete(&player).Error; err != nil {
		http.Error(w, "Failed to delete player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

func (s *PlayerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// get id param
	playerIdParam := chi.URLParam(r, "id")
	if playerIdParam == "" {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	// check if player exists
	var player models.Player
	if err := db.DB.First(&player, playerIdParam).Error; err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}
