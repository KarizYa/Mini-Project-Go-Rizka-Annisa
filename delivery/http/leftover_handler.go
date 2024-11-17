package http

import (
	"mini-project/models"
	"mini-project/usecases"
	"mini-project/helper" 
	"net/http"
	"strconv"
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
		return c.JSON(http.StatusUnauthorized, helper.WrapResponse("User not authenticated", 401, "error", nil))
	}

	var leftover models.Leftover
	if err := c.Bind(&leftover); err != nil {
		return c.JSON(http.StatusBadRequest, helper.WrapResponse("Invalid input", 400, "error", nil))
	}

	leftover.UserID = userID 
	if err := h.Usecase.CreateLeftover(&leftover); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WrapResponse("Failed to create leftover", 500, "error", nil))
	}

	return c.JSON(http.StatusCreated, helper.WrapResponse("Leftover created successfully", 201, "success", leftover))
}

func (h *LeftoverHandler) GetAllLeftovers(c echo.Context) error {
	userID, ok := c.Get("userID").(uint) 
	if !ok {
		return c.JSON(http.StatusUnauthorized, helper.WrapResponse("User not authenticated", 401, "error", nil))
	}

	leftovers, err := h.Usecase.GetAllLeftovers(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WrapResponse("Failed to fetch leftovers", 500, "error", nil))
	}

	return c.JSON(http.StatusOK, helper.WrapResponse("Successfully fetched leftovers", 200, "success", leftovers))
}

func (h *LeftoverHandler) GetLeftoverByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WrapResponse("Invalid ID", 400, "error", nil))
	}

	leftover, err := h.Usecase.GetLeftoverByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.WrapResponse("Leftover not found", 404, "error", nil))
	}

	return c.JSON(http.StatusOK, helper.WrapResponse("Successfully fetched leftover", 200, "success", leftover))
}

func (h *LeftoverHandler) UpdateLeftover(c echo.Context) error {
	var leftover models.Leftover
	if err := c.Bind(&leftover); err != nil {
		return c.JSON(http.StatusBadRequest, helper.WrapResponse("Invalid input", 400, "error", nil))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WrapResponse("Invalid ID", 400, "error", nil))
	}

	existingLeftover, err := h.Usecase.GetLeftoverByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.WrapResponse("Leftover not found", 404, "error", nil))
	}

	existingLeftover.Name = leftover.Name
	existingLeftover.Quantity = leftover.Quantity
	existingLeftover.Unit = leftover.Unit
	existingLeftover.ExpiryDate = leftover.ExpiryDate

	if err := h.Usecase.UpdateLeftover(&existingLeftover); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WrapResponse("Failed to update leftover", 500, "error", nil))
	}

	return c.JSON(http.StatusOK, helper.WrapResponse("Leftover updated successfully", 200, "success", existingLeftover))
}


func (h *LeftoverHandler) DeleteLeftover(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WrapResponse("Invalid ID", 400, "error", nil))
	}

	if err := h.Usecase.DeleteLeftover(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WrapResponse("Failed to delete leftover", 500, "error", nil))
	}

	return c.JSON(http.StatusOK, helper.WrapResponse("Leftover deleted successfully", 200, "success", nil))
}
