package http

import (
	"mini-project/usecases"
	"mini-project/helper" 
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
		return c.JSON(http.StatusInternalServerError, helper.WrapResponse("Failed to retrieve leaderboards", 500, "error", nil))
	}

	return c.JSON(http.StatusOK, helper.WrapResponse("Successfully retrieved all leaderboards", 200, "success", leaderboards))
}

func (h *LeaderboardHandler) GetLeaderboardByID(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WrapResponse("Invalid ID", 400, "error", nil))
	}

	leaderboard, err := h.leaderboardUsecase.GetLeaderboardByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WrapResponse("Failed to retrieve leaderboard", 500, "error", nil))
	}

	return c.JSON(http.StatusOK, helper.WrapResponse("Successfully retrieved leaderboard", 200, "success", leaderboard))
}

func (h *LeaderboardHandler) AddToLeaderboard(c echo.Context) error {
	var req struct {
		UserID uint `json:"user_id"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.WrapResponse("Invalid input", 400, "error", nil))
	}

	if req.UserID == 0 {
		return c.JSON(http.StatusBadRequest, helper.WrapResponse("User ID is required", 400, "error", nil))
	}

	createdLeaderboard, err := h.leaderboardUsecase.CreateLeaderboard(req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WrapResponse("Failed to create leaderboard", 500, "error", nil))
	}

	return c.JSON(http.StatusCreated, helper.WrapResponse("Successfully added to leaderboard", 201, "success", createdLeaderboard))
}
