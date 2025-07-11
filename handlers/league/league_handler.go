package league

import (
	"encoding/json"
	"leaguies_backend/internal/db"
	"leaguies_backend/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"leaguies_backend/internal/db/league"
)

type LeagueHandler struct {
	Store league.LeagueStoreInterface
}

type CreateLeagueRequest struct {
	Name    string `json:"name"`
	SportId uint   `json:"sport_id"`
}

type UpdateLeagueRequest struct {
	Name    *string `json:"name"`
	SportId *uint   `json:"sport_id"`
}

func NewLeagueHandler(store league.LeagueStoreInterface) *LeagueHandler {
	return &LeagueHandler{
		Store: store,
	}
}

func (h *LeagueHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateLeagueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.SportId == 0 {
		http.Error(w, `{"error":"Missing required fields: name and sportId"}`, http.StatusBadRequest)
		return
	}

	league := models.League{
		Name:    req.Name,
		SportID: req.SportId,
	}

	if err := h.Store.Create(&league); err != nil {
		http.Error(w, "Failed to create league", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(league)
}

func (h *LeagueHandler) Update(w http.ResponseWriter, r *http.Request) {
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

func (h *LeagueHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

func (h *LeagueHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

func (h *LeagueHandler) List(w http.ResponseWriter, r *http.Request) {
	var leagues []models.League
	if err := db.DB.Find(&leagues).Error; err != nil {
		http.Error(w, "Failed to fetch leagues", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(leagues)
}
