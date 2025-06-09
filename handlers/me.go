package handlers

import (
	"encoding/json"
	"net/http"

	"leaguies_backend/internal/db"
	"leaguies_backend/middleware"
	"leaguies_backend/models"
)

func Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
