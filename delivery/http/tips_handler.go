package http

import (
    "mini-project/models"
    "mini-project/usecases"
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

    return c.JSON(http.StatusOK, tips)
}

func (h *TipsHandler) GetTipsByLeftover(c echo.Context) error {
    ingredient := c.QueryParam("ingredient")
    tips, err := h.TipsUsecase.GetTipsByLeftover(ingredient) 
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch tips")
    }

    return c.JSON(http.StatusOK, tips)
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

    if err := h.TipsUsecase.CreateTips(req); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create tip")
    }

    return c.NoContent(http.StatusCreated)
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

    if err := h.TipsUsecase.UpdateTips(req); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update tip")
    }

    return c.NoContent(http.StatusOK)
}

func (h *TipsHandler) DeleteTips(c echo.Context) error {
    userID, ok := c.Get("userID").(uint) 
    if !ok {
        return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
    }

    idStr := c.QueryParam("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
    }

    if err := h.TipsUsecase.DeleteTips(userID, uint(id)); 
    err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete tip")
    }

    return c.NoContent(http.StatusNoContent)
}
