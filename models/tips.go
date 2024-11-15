// models/tips.go
package models

import (
	"strings"
	"time"
)

type Tips struct {
    ID        uint64   `json:"id"`
    UserID    uint64   `json:"user_id"`  // Gunakan uint64 di sini
    Title     string   `json:"title"`
    Content   string   `json:"content"`
    Leftovers string   `json:"leftovers"` // Data disimpan sebagai string di database
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
