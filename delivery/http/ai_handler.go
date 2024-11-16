package http

import (
	"mini-project/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SuggestionHandler struct {
	SuggestionUseCase *usecases.SuggestionUseCase
}

func NewSuggestionHandler(suggestionUseCase *usecases.SuggestionUseCase) *SuggestionHandler {
	return &SuggestionHandler{
		SuggestionUseCase: suggestionUseCase,
	}
}

func (h *SuggestionHandler) GetSuggestionsHandler(c echo.Context) error {
	leftover := c.QueryParam("leftover")
	if leftover == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "leftover parameter is required",
		})
	}

	suggestion, err := h.SuggestionUseCase.GetSuggestion(leftover)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, suggestion)
}
