package http

import (
    "mini-project/usecases"
    "mini-project/helper" 
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
        return c.JSON(http.StatusBadRequest, helper.WrapResponse("meal_name query parameter is required", 400, "error", nil))
    }

    recipes, err := h.recipeUsecase.GetRecipesByName(mealName)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, helper.WrapResponse(err.Error(), 500, "error", nil))
    }

    return c.JSON(http.StatusOK, helper.WrapResponse("Successfully fetched recipes", 200, "success", recipes))
}
