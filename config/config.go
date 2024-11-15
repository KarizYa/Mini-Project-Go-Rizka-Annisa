package config

import (
	"fmt"
	"log"
	"os"
	"mini-project/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Load .env file
func LoadEnvVariables() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }
}

// InitDB initializes the database connection
func InitDB() (*gorm.DB, error) {
    LoadEnvVariables()

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    db.AutoMigrate(&models.User{}, &models.Leftover{}, &models.Recipe{}, &models.Tips{}) 
    return db, nil
}
