package models

type Leaderboard struct {
    ID      uint `json:"leaderboard_id" gorm:"primaryKey"`
    UserID  uint `json:"user_id"`
    Points  int  `json:"points"`
}
