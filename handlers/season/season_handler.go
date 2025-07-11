package season

import (
	"encoding/json"
	"leaguies_backend/internal/db"
	"leaguies_backend/models"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"leaguies_backend/internal/db/season"
)

type SeasonHandler struct {
	store season.SeasonStoreInterface
}

type CreateSeasonRequest struct {
	Name      string    `json:"name"`
	LeagueID  uint      `json:"league_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type UpdateSeasonRequest struct {
	Name      *string    `json:"name"`
	LeagueID  *uint      `json:"league_id"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

func NewSeasonHandler(store season.SeasonStoreInterface) *SeasonHandler {
	return &SeasonHandler{
		store: store,
	}
}

func (h *SeasonHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	if err := h.store.Create(&season); err != nil {
		http.Error(w, "Failed to create season", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(season)
}

func (h *SeasonHandler) Update(w http.ResponseWriter, r *http.Request) {
	// get id param
	seasonIdParam := chi.URLParam(r, "id")
	if seasonIdParam == "" {
		http.Error(w, "Invalid season ID", http.StatusBadRequest)
		return
	}

	// decode request
	var req UpdateSeasonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// check if season exists
	var season models.Season
	if err := db.DB.First(&season, seasonIdParam).Error; err != nil {
		http.Error(w, "Season not found", http.StatusNotFound)
		return
	}

	// update
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

	if err := h.store.Update(&season); err != nil {
		http.Error(w, "Failed to update season", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(season)
}	

func (h *SeasonHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// get id param
	seasonIdParam := chi.URLParam(r, "id")
	if seasonIdParam == "" {
		http.Error(w, "Invalid season ID", http.StatusBadRequest)
		return
	}

	// check if season exists
	var season models.Season
	if err := db.DB.First(&season, seasonIdParam).Error; err != nil {
		http.Error(w, "Season not found", http.StatusNotFound)
		return
	}

	// delete
	if err := h.store.Delete(&season); err != nil {
		http.Error(w, "Failed to delete season", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(season)
}

func (h *SeasonHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// get id param
	seasonIdParam := chi.URLParam(r, "id")
	if seasonIdParam == "" {
		http.Error(w, "Invalid season ID", http.StatusBadRequest)
		return
	}

	// check if season exists
	var season models.Season
	if err := db.DB.First(&season, seasonIdParam).Error; err != nil {
		http.Error(w, "Season not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(season)
}

func (h *SeasonHandler) List(w http.ResponseWriter, r *http.Request) {
	var seasons []models.Season
	if err := db.DB.Find(&seasons).Error; err != nil {
		http.Error(w, "Failed to fetch seasons", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(seasons)
}