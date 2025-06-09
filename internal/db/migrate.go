package db

import (
    "log"

    "leaguies_backend/models"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    err := db.AutoMigrate(
        &models.User{},
		// add other models here
    )
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
    log.Println("Database migration completed")
}
