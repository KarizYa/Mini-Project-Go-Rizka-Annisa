package models

import (
    "strings"
    "time"
)

type Tips struct {
    ID        uint     `json:"id" gorm:"primaryKey;autoIncrement"`       
    UserID    uint     `json:"user_id"`  // Ubah menjadi uint
    Title     string   `json:"title"`
    Content   string   `json:"content"`
    Leftovers string   `json:"leftovers"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Fungsi untuk mengonversi Leftovers dari string ke slice string
func (t *Tips) GetLeftoversSlice() []string {
    if t.Leftovers == "" {
        return []string{}
    }
    return strings.Split(t.Leftovers, ",")
}

// Fungsi untuk mengonversi Leftovers dari slice string ke string
func (t *Tips) SetLeftoversSlice(leftovers []string) {
    t.Leftovers = strings.Join(leftovers, ",")
}
