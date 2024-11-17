package http

import (
	"mini-project/models"
	"mini-project/usecases"
	"mini-project/helper" 
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TipsHandler struct {
	TipsUsecase *usecases.TipsUsecase
}

func NewTipsHandler(tipsUsecase *usecases.TipsUsecase) *TipsHandler {
	return &TipsHandler{
		TipsUsecase: tipsUsecase,
	}
}

func (h *TipsHandler) GetAllTips(c echo.Context) error {
	tips, err := h.TipsUsecase.GetAllTips()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch tips")
	}

	return c.JSON(http.StatusOK, helper.WrapResponse("Successfully fetched all tips", 200, "success", tips))
}

func (h *TipsHandler) GetTipsByLeftover(c echo.Context) error {
    leftover := c.QueryParam("leftovers") 
    if leftover == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "leftovers parameter is required")
    }

    tips, err := h.TipsUsecase.GetTipsByLeftover(leftover)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch tips")
    }

    return c.JSON(http.StatusOK, helper.WrapResponse("Successfully fetched tips for leftover", 200, "success", tips))
}


func (h *TipsHandler) CreateTips(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	var req models.Tips
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	req.UserID = userID

	err := h.TipsUsecase.CreateTips(req) 
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create tip")
	}

	return c.JSON(http.StatusCreated, helper.WrapResponse("Successfully created tip", 201, "success", nil))
}

func (h *TipsHandler) UpdateTips(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	var req models.Tips
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	req.UserID = userID

	err := h.TipsUsecase.UpdateTips(req) 
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update tip")
	}

	return c.JSON(http.StatusOK, helper.WrapResponse("Successfully updated tip", 200, "success", nil))
}

func (h *TipsHandler) DeleteTips(c echo.Context) error {
    userID, ok := c.Get("userID").(uint)
    if !ok {
        return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
    }

    idStr := c.Param("id") 
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
    }

    err = h.TipsUsecase.DeleteTips(userID, uint(id))
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete tip")
    }

    return c.JSON(http.StatusOK, helper.WrapResponse("Successfully deleted tip", 200, "success", map[string]interface{}{"id": id}))
}

