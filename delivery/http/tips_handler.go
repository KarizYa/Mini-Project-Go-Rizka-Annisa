// handlers/tips_handler.go
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

// Constructor untuk TipsHandler
func NewTipsHandler(tipsUsecase *usecases.TipsUsecase) *TipsHandler {
    return &TipsHandler{
        TipsUsecase: tipsUsecase,
    }
}

// Fungsi untuk mendapatkan semua tips
func (h *TipsHandler) GetAllTips(c echo.Context) error {
    // Mengambil user_id dari konteks
    userID, ok := c.Get("userID").(uint64) // Mengubah menjadi uint64
    if !ok {
        return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
    }

    tips, err := h.TipsUsecase.GetAllTips(userID) // Menggunakan uint64
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch tips")
    }

    return c.JSON(http.StatusOK, tips)
}

// Fungsi untuk mendapatkan tips berdasarkan sisa makanan
func (h *TipsHandler) GetTipsByLeftover(c echo.Context) error {
    // Mengambil user_id dari konteks
    userID, ok := c.Get("userID").(uint64) // Mengubah menjadi uint64
    if !ok {
        return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
    }

    ingredient := c.QueryParam("ingredient")
    tips, err := h.TipsUsecase.GetTipsByLeftover(userID, ingredient) // Menggunakan uint64
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch tips")
    }

    return c.JSON(http.StatusOK, tips)
}

// Fungsi untuk menambahkan tips baru
func (h *TipsHandler) CreateTips(c echo.Context) error {
    // Mengambil user_id dari konteks
    userID, ok := c.Get("userID").(uint64) // Mengubah menjadi uint64
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

// Fungsi untuk memperbarui tips
func (h *TipsHandler) UpdateTips(c echo.Context) error {
    // Mengambil user_id dari konteks
    userID, ok := c.Get("userID").(uint64) // Mengubah menjadi uint64
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

// Fungsi untuk menghapus tips
func (h *TipsHandler) DeleteTips(c echo.Context) error {
    // Mengambil user_id dari konteks
    userID, ok := c.Get("userID").(uint64) // Mengubah menjadi uint64
    if !ok {
        return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
    }

    idStr := c.QueryParam("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
    }

    if err := h.TipsUsecase.DeleteTips(userID, uint(id)); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete tip")
    }

    return c.NoContent(http.StatusNoContent)
}
