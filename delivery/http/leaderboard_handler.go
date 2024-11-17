package http

import (
	"mini-project/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LeaderboardHandler struct {
	leaderboardUsecase usecases.LeaderboardUsecase
}

func NewLeaderboardHandler(leaderboardUsecase usecases.LeaderboardUsecase) *LeaderboardHandler {
	return &LeaderboardHandler{
		leaderboardUsecase: leaderboardUsecase,
	}
}

func (h *LeaderboardHandler) GetAllLeaderboards(c echo.Context) error {
	leaderboards, err := h.leaderboardUsecase.GetAllLeaderboards()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, leaderboards)
}

func (h *LeaderboardHandler) GetLeaderboardByID(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	leaderboard, err := h.leaderboardUsecase.GetLeaderboardByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, leaderboard)
}

func (h *LeaderboardHandler) AddToLeaderboard(c echo.Context) error {
	var req struct {
		UserID uint `json:"user_id"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if req.UserID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	createdLeaderboard, err := h.leaderboardUsecase.CreateLeaderboard(req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create leaderboard"})
	}

	return c.JSON(http.StatusCreated, createdLeaderboard)
}


