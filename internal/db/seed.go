// internal/db/seed.go
package db

import (
	"leaguies_backend/models"
	"log"
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

func SeedSports() {
	sports := []models.Sport{
		{Name: "Football"},
		{Name: "Basketball"},
		{Name: "Volleyball"},
		{Name: "Baseball"},
		{Name: "Soccer"},
	}

	for _, sport := range sports {
		var existing models.Sport
		if err := DB.Where("name = ?", sport.Name).First(&existing).Error; err != nil {
			if err := DB.Create(&sport).Error; err != nil {
				log.Printf("Failed to seed sport %s: %v", sport.Name, err)
			}
		}
	}
}
