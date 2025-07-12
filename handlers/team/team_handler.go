package team

import (
	"encoding/json"
	"leaguies_backend/internal/db/team"
	"leaguies_backend/internal/utils"
	"leaguies_backend/models"
	"net/http"
)

type TeamHandler struct {
	store team.TeamStoreInterface
}
type CreateTeamRequest struct {
	Name     string `json:"name"`
	LeagueId *uint  `json:"league_id"`
}

type UpdateTeamRequest struct {
	Name     *string `json:"name"`
	LeagueId *uint   `json:"league_id"`
}

func NewTeamHandler(store team.TeamStoreInterface) *TeamHandler {
	return &TeamHandler{
		store: store,
	}
}

func (h *TeamHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateTeamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	team := models.Team{
		Name:     req.Name,
		LeagueId: req.LeagueId,
	}

	if err := h.store.Create(&team); err != nil {
		http.Error(w, "Failed to create team", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}

func (h *TeamHandler) Update(w http.ResponseWriter, r *http.Request) {
	uint, error := utils.ParseUintParam(r, "id")
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	team, err := h.store.GetByID(uint)
	if err != nil {
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

	if err := h.store.Update(team); err != nil {
		http.Error(w, "Failed to update team", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}

func (h *TeamHandler) Delete(w http.ResponseWriter, r *http.Request) {
	uint, error := utils.ParseUintParam(r, "id")
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	team, err := h.store.GetByID(uint)
	if err != nil {
		http.Error(w, "Team not found", http.StatusNotFound)
		return
	}

	if err := h.store.Delete(team); err != nil {
		http.Error(w, "Failed to delete team", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TeamHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	uint, error := utils.ParseUintParam(r, "id")
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	team, err := h.store.GetByID(uint)
	if err != nil {
		http.Error(w, "Team not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}

func (h *TeamHandler) List(w http.ResponseWriter, r *http.Request) {
	teams, err := h.store.List()
	if err != nil {
		http.Error(w, "Failed to fetch teams", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}
