package http

import (
    "mini-project/usecases"
    "net/http"
    "github.com/labstack/echo/v4"
)

type RecipeHandler struct {
    recipeUsecase *usecases.RecipeUsecase
}

// Fungsi untuk membuat instance RecipeHandler baru
func NewRecipeHandler(recipeUsecase *usecases.RecipeUsecase) *RecipeHandler {
    return &RecipeHandler{recipeUsecase: recipeUsecase}
}

// Fungsi handler untuk pencarian resep berdasarkan nama makanan
func (h *RecipeHandler) SearchRecipesHandler(c echo.Context) error {
    mealName := c.QueryParam("meal_name")
    if mealName == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "meal_name query parameter is required"})
    }

    // Mengambil resep berdasarkan nama makanan menggunakan usecase
    recipes, err := h.recipeUsecase.GetRecipesByName(mealName)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    // Mengembalikan resep yang ditemukan dalam format JSON
    return c.JSON(http.StatusOK, recipes)
}
