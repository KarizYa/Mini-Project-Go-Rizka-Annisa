package repositories

import (
	"mini-project/models"
	"gorm.io/gorm"
)

type LeaderboardRepository interface {
	GetAll() ([]models.Leaderboard, error)
	GetByID(id uint) (models.Leaderboard, error)
	GetByUserID(userID uint) (models.Leaderboard, error)
	CreateLeaderboard(leaderboard *models.Leaderboard) error
}

type leaderboardRepository struct {
	DB *gorm.DB
}

func NewLeaderboardRepository(DB *gorm.DB) LeaderboardRepository {
	return &leaderboardRepository{DB: DB}
}

func (repo *leaderboardRepository) GetAll() ([]models.Leaderboard, error) {
	var leaderboards []models.Leaderboard
	err := repo.DB.Order("points desc").Find(&leaderboards).Error
	return leaderboards, err
}

func (repo *leaderboardRepository) GetByID(id uint) (models.Leaderboard, error) {
	var leaderboard models.Leaderboard
	err := repo.DB.First(&leaderboard, id).Error
	return leaderboard, err
}

func (repo *leaderboardRepository) GetByUserID(userID uint) (models.Leaderboard, error) {
	var leaderboard models.Leaderboard
	err := repo.DB.Where("user_id = ?", userID).First(&leaderboard).Error
	if err != nil {
		return models.Leaderboard{}, err
	}
	return leaderboard, nil
}

func (repo *leaderboardRepository) CreateLeaderboard(leaderboard *models.Leaderboard) error {
	return repo.DB.Create(leaderboard).Error
}
