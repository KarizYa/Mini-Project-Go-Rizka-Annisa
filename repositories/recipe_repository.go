package repositories

import "mini-project/models"

type RecipeRepository interface {
    FindByID(id string) (models.Recipe, error)
}

type recipeRepository struct{}

func NewRecipeRepository() RecipeRepository {
    return &recipeRepository{}
}

func (r *recipeRepository) FindByID(id string) (models.Recipe, error) {
    return models.Recipe{}, nil
}
