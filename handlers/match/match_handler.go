package match

import (
	"encoding/json"
	"leaguies_backend/internal/db/match"
	"leaguies_backend/internal/db/team"
	"leaguies_backend/internal/utils"
	"leaguies_backend/models"
	"net/http"
	"time"
)

type MatchHandler struct {
	store     match.MatchStoreInterface
	TeamStore team.TeamStoreInterface
}

type CreateMatchRequest struct {
	HomeTeamID uint `json:"home_team_id"`
	AwayTeamID uint `json:"away_team_id"`
	SeasonID   uint `json:"season_id"`
}

type UpdateMatchRequest struct {
	ID          uint       `json:"id"`
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

func NewMatchHandler(store match.MatchStoreInterface, teamStore team.TeamStoreInterface) *MatchHandler {
	return &MatchHandler{
		store:     store,
		TeamStore: teamStore,
	}
}

func (h *MatchHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if _, err := h.TeamStore.GetByID(req.HomeTeamID); err != nil {
		http.Error(w, "Home team not found", http.StatusNotFound)
		return
	}

	if _, err := h.TeamStore.GetByID(req.AwayTeamID); err != nil {
		http.Error(w, "Away team not found", http.StatusNotFound)
		return
	}

	match := models.Match{
		Name:        "Match",
		SeasonID:    &req.SeasonID,
		HomeTeamID:  &req.HomeTeamID,
		AwayTeamID:  &req.AwayTeamID,
		HomeScore:   nil,
		AwayScore:   nil,
		Status:      nil,
		Location:    nil,
		ScheduledAt: nil,
	}

	if err := h.store.Create(&match); err != nil {
		http.Error(w, "Failed to create match", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

func (h *MatchHandler) Update(w http.ResponseWriter, r *http.Request) {
	uint, error := utils.ParseUintParam(r, "id")
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	match, err := h.store.GetByID(uint)
	if err != nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	var req UpdateMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.SeasonID != nil {
		match.SeasonID = req.SeasonID
	}

	if req.HomeTeamID != nil {
		match.HomeTeamID = req.HomeTeamID
	}

	if req.AwayTeamID != nil {
		match.AwayTeamID = req.AwayTeamID
	}

	if req.HomeScore != nil {
		match.HomeScore = req.HomeScore
	}

	if req.AwayScore != nil {
		match.AwayScore = req.AwayScore
	}

	if req.Status != nil {
		match.Status = req.Status
	}

	if req.Location != nil {
		match.Location = req.Location
	}

	if req.ScheduledAt != nil {
		match.ScheduledAt = req.ScheduledAt
	}

	if err := h.store.Update(match); err != nil {
		http.Error(w, "Failed to update match", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)
}

func (h *MatchHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	uint, error := utils.ParseUintParam(r, "id")
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	match, err := h.store.GetByID(uint)
	if err != nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)
}

func (h *MatchHandler) List(w http.ResponseWriter, r *http.Request) {
	teams, err := h.store.List()
	if err != nil {
		http.Error(w, "Failed to fetch matches", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}

func (h *MatchHandler) Delete(w http.ResponseWriter, r *http.Request) {
	uint, error := utils.ParseUintParam(r, "id")
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	match, err := h.store.GetByID(uint)
	if err != nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	if err := h.store.Delete(match); err != nil {
		http.Error(w, "Failed to delete match", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
