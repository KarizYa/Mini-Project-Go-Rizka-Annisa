package usecases

import (
	"mini-project/infrastructures/external"
	"mini-project/models"
)

type RecipeUsecase struct {
    recipeAPI *external.RecipeAPI
}

// Fungsi untuk membuat instance RecipeUsecase baru
func NewRecipeUsecase(recipeAPI *external.RecipeAPI) *RecipeUsecase {
    return &RecipeUsecase{recipeAPI: recipeAPI}
}

// Fungsi untuk mendapatkan resep berdasarkan nama makanan
func (u *RecipeUsecase) GetRecipesByName(mealName string) ([]models.Recipe, error) {
    return u.recipeAPI.GetRecipesByName(mealName)
}
