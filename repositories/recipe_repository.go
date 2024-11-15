package repositories

import "mini-project/models"

// Hapus repository untuk data lokal karena kita akan mengakses API eksternal.
type RecipeRepository interface {
    FindByID(id string) (models.Recipe, error)
}

type recipeRepository struct{}

func NewRecipeRepository() RecipeRepository {
    return &recipeRepository{}
}

func (r *recipeRepository) FindByID(id string) (models.Recipe, error) {
    // Akan diproses oleh external API
    return models.Recipe{}, nil
}
