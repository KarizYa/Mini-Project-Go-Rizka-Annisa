package usecases

import (
	"errors"
	"mini-project/models"
	"mini-project/repositories"
)

type LeaderboardUsecase interface {
	GetAllLeaderboards() ([]models.Leaderboard, error)
	GetLeaderboardByID(id uint) (models.Leaderboard, error)
	CreateLeaderboard(userID uint) (models.Leaderboard, error)
}

type leaderboardUsecase struct {
	leaderboardRepo repositories.LeaderboardRepository
	userRepo        repositories.UserRepository
}

func NewLeaderboardUsecase(
	leaderboardRepo repositories.LeaderboardRepository,
	userRepo repositories.UserRepository,
) LeaderboardUsecase {
	return &leaderboardUsecase{
		leaderboardRepo: leaderboardRepo,
		userRepo:        userRepo,
	}
}

func (uc *leaderboardUsecase) GetAllLeaderboards() ([]models.Leaderboard, error) {
	return uc.leaderboardRepo.GetAll()
}

func (uc *leaderboardUsecase) GetLeaderboardByID(id uint) (models.Leaderboard, error) {
	return uc.leaderboardRepo.GetByID(id)
}

func (uc *leaderboardUsecase) CreateLeaderboard(userID uint) (models.Leaderboard, error) {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return models.Leaderboard{}, errors.New("user not found")
	}

	existingLeaderboard, _ := uc.leaderboardRepo.GetByUserID(userID)
	if existingLeaderboard.ID != 0 {
		return models.Leaderboard{}, errors.New("user already exists in the leaderboard")
	}

	newLeaderboard := models.Leaderboard{
		UserID: user.ID,
		Points: user.Points,
	}

	err = uc.leaderboardRepo.CreateLeaderboard(&newLeaderboard)
	if err != nil {
		return models.Leaderboard{}, errors.New("failed to create leaderboard")
	}

	return newLeaderboard, nil
}

