package http

import (
    "mini-project/usecases"
    "net/http"
    "github.com/labstack/echo/v4"
)

type RecipeHandler struct {
    recipeUsecase *usecases.RecipeUsecase
}

func NewRecipeHandler(recipeUsecase *usecases.RecipeUsecase) *RecipeHandler {
    return &RecipeHandler{recipeUsecase: recipeUsecase}
}

func (h *RecipeHandler) SearchRecipesHandler(c echo.Context) error {
    mealName := c.QueryParam("meal_name")
    if mealName == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "meal_name query parameter is required"})
    }

    recipes, err := h.recipeUsecase.GetRecipesByName(mealName)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, recipes)
}
