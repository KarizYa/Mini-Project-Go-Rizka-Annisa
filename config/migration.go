package config

import (
	"log"
	"mini-project/models"

	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
    err := db.AutoMigrate(&models.User{}, &models.Leftover{}, &models.Recipe{}, &models.Tips{}, &models.Leaderboard{})
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }
}
