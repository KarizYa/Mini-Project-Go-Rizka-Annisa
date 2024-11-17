package usecases_test

import (
	"errors"
	"mini-project/mocks"
	"mini-project/models"
	"mini-project/usecases"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateLeaderboard_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLeaderboardRepo := mocks.NewMockLeaderboardRepository(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	usecase := usecases.NewLeaderboardUsecase(mockLeaderboardRepo, mockUserRepo)

	mockUserRepo.EXPECT().GetByID(uint(1)).Return(models.User{}, errors.New("user not found"))

	_, err := usecase.CreateLeaderboard(1)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}

func TestCreateLeaderboard_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLeaderboardRepo := mocks.NewMockLeaderboardRepository(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	usecase := usecases.NewLeaderboardUsecase(mockLeaderboardRepo, mockUserRepo)

	mockUser := models.User{
		Email:    "test@example.com",
		Password: "securepassword",
		Points:   100,
	}

	mockUserRepo.EXPECT().GetByID(uint(1)).Return(mockUser, nil)

	mockLeaderboardRepo.EXPECT().GetByUserID(uint(1)).Return(models.Leaderboard{}, nil) 

	mockLeaderboardRepo.EXPECT().CreateLeaderboard(gomock.Any()).Return(nil).Times(1) 

	result, err := usecase.CreateLeaderboard(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockUser.Points, result.Points)
}

func TestCreateLeaderboard_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockLeaderboardRepo := mocks.NewMockLeaderboardRepository(ctrl)
    mockUserRepo := mocks.NewMockUserRepository(ctrl)

    usecase := usecases.NewLeaderboardUsecase(mockLeaderboardRepo, mockUserRepo)

    mockUserRepo.EXPECT().GetByID(uint(1)).Return(models.User{}, errors.New("user not found"))

    _, err := usecase.CreateLeaderboard(1)
    assert.Error(t, err)
    assert.Equal(t, "user not found", err.Error())
}

