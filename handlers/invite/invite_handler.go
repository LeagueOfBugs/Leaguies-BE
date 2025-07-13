package invite

import (
	"encoding/json"
	"leaguies_backend/internal/db/invite"
	"leaguies_backend/internal/utils"
	"leaguies_backend/models"
	"net/http"
	"time"
)

type InviteHandler struct {
	store invite.InviteStoreInterface
}

type CreateInviteRequest struct {
	Type      string `json:"type"`
	FromID    uint   `json:"from_id"`
	ToID      uint   `json:"to_id"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateInviteRequest struct {
	Type      *string `json:"type"`
	FromID    *uint   `json:"from_id"`
	ToID      *uint   `json:"to_id"`
	Status    *string `json:"status"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewInviteHandler(store invite.InviteStoreInterface) *InviteHandler {
	return &InviteHandler{
		store: store,
	}
}

func (h *InviteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateInviteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	invite := models.Invite{
		Type:      req.Type,
		FromID:    req.FromID,
		ToID:      req.ToID,
		Status:    req.Status,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}

	if err := h.store.Create(&invite); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invite)
}

func (h *InviteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	uint, error := utils.ParseUintParam(r, "id")
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	invite, err := h.store.GetByID(uint)
	if err != nil {
		http.Error(w, "Invite not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invite)
}

func (h *InviteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	uint, error := utils.ParseUintParam(r, "id")
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	invite, err := h.store.GetByID(uint)
	if err != nil {
		http.Error(w, "Invite not found", http.StatusNotFound)
		return
	}

	if err := h.store.Delete(invite); err != nil {
		http.Error(w, "Failed to delete invite", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *InviteHandler) List(w http.ResponseWriter, r *http.Request) {
	invites, err := h.store.List()
	if err != nil {
		http.Error(w, "Failed to fetch invites", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invites)
}
