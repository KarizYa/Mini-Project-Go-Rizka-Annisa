package usecases

import (
	"mini-project/infrastructures/external"
	"mini-project/models"
)

type RecipeUsecase struct {
    recipeAPI *external.RecipeAPI
}

func NewRecipeUsecase(recipeAPI *external.RecipeAPI) *RecipeUsecase {
    return &RecipeUsecase{recipeAPI: recipeAPI}
}

func (u *RecipeUsecase) GetRecipesByName(mealName string) ([]models.Recipe, error) {
    return u.recipeAPI.GetRecipesByName(mealName)
}
