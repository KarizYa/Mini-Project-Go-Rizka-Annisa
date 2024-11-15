package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Password   string    `json:"password"`
    Email      string    `json:"email" gorm:"unique"`
    Points     int       `json:"points"`
    CreatedAt  time.Time `json:"created_at"`
    DeletedAt  time.Time `json:"deleted_at"`
}

type Credentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
