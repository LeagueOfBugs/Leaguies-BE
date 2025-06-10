package handlers

import (
	"encoding/json"
	"net/http"

	"leaguies_backend/internal/db"
	"leaguies_backend/models"
	"leaguies_backend/utils"
)

type RegisterRequest struct {
    Email    string   `json:"email"`
    Password string   `json:"password"`
    FullName string   `json:"full_name"`
    Roles    []string `json:"roles"`
}

func Register(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Check if user already exists
    var existingUser models.User
    if err := db.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
        http.Error(w, "User already exists", http.StatusBadRequest)
        return
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // Fetch roles from DB based on names provided (or assign default role)
    var roles []models.Role
    if len(req.Roles) == 0 {
        // If no roles specified, assign default "user" role
        var defaultRole models.Role
        if err := db.DB.Where("name = ?", "user").First(&defaultRole).Error; err != nil {
            http.Error(w, "Default role not found", http.StatusInternalServerError)
            return
        }
        roles = append(roles, defaultRole)
    } else {
        if err := db.DB.Where("name IN ?", req.Roles).Find(&roles).Error; err != nil {
            http.Error(w, "Roles not found", http.StatusBadRequest)
            return
        }
    }

    // Create user
    user := models.User{
        Email:    req.Email,
        Password: hashedPassword,
        FullName: req.FullName,
        Roles:    roles,
    }

    // Save user and related roles in one transaction
    if err := db.DB.Create(&user).Error; err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    // Check if user has role "player"
    var hasPlayerRole bool
    for _, role := range roles {
        if role.Name == "player" {
            hasPlayerRole = true
            break
        }
    }

    // If player role assigned, create empty Player linked to this user
    if hasPlayerRole {
        player := models.Player{
            UserID: user.ID,
        }
        if err := db.DB.Create(&player).Error; err != nil {
            // Optional: rollback user creation or log error, depends on your logic
            http.Error(w, "Failed to create player profile", http.StatusInternalServerError)
            return
        }
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
