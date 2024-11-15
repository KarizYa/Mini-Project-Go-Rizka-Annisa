package models

import "time"

type Leftover struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    UserID      uint      `json:"user_id"`
    Name        string    `json:"name"`
    Quantity    float64   `json:"quantity"`
    Unit        string    `json:"unit"`
    ExpiryDate  time.Time `json:"expiry_date"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

