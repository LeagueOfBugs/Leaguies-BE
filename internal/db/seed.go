// internal/db/seed.go
package db

import (
	"log"
	"leaguies_backend/models"
)

func SeedRoles() {
	roles := []models.Role{
		{Name: "user"},
		{Name: "player"},
		{Name: "coach"},
		{Name: "referee"},
	}

	for _, role := range roles {
		var existing models.Role
		if err := DB.Where("name = ?", role.Name).First(&existing).Error; err != nil {
			if err := DB.Create(&role).Error; err != nil {
				log.Printf("Failed to seed role %s: %v", role.Name, err)
			}
		}
	}
}
