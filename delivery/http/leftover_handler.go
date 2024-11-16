package http

import (
	"net/http"
	"strconv"

	"mini-project/models"
	"mini-project/usecases"

	"github.com/labstack/echo/v4"
)

type LeftoverHandler struct {
	Usecase usecases.LeftoverUsecase
}

func NewLeftoverHandler(u usecases.LeftoverUsecase) *LeftoverHandler {
	return &LeftoverHandler{Usecase: u}
}

func (h *LeftoverHandler) CreateLeftover(c echo.Context) error {
	userID, ok := c.Get("userID").(uint) 
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User not authenticated"})
	}

	var leftover models.Leftover
	if err := c.Bind(&leftover); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	leftover.UserID = userID 
	if err := h.Usecase.CreateLeftover(&leftover); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create leftover"})
	}

	return c.JSON(http.StatusCreated, leftover)
}

func (h *LeftoverHandler) GetAllLeftovers(c echo.Context) error {
	userID, ok := c.Get("userID").(uint) 
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User not authenticated"})
	}

	leftovers, err := h.Usecase.GetAllLeftovers(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch leftovers"})
	}

	return c.JSON(http.StatusOK, leftovers)
}

func (h *LeftoverHandler) GetLeftoverByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	leftover, err := h.Usecase.GetLeftoverByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Leftover not found"})
	}

	return c.JSON(http.StatusOK, leftover)
}

func (h *LeftoverHandler) UpdateLeftover(c echo.Context) error {
	var leftover models.Leftover
	if err := c.Bind(&leftover); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	if err := h.Usecase.UpdateLeftover(&leftover); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update leftover"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Leftover updated successfully"})
}

func (h *LeftoverHandler) DeleteLeftover(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	if err := h.Usecase.DeleteLeftover(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete leftover"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Leftover deleted successfully"})
}
