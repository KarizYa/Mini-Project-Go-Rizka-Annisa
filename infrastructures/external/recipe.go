package external

import (
    "encoding/json"
    "fmt"
    "mini-project/models"
    "net/http"
)

type RecipeAPI struct {
    BaseURL string
}

// Fungsi untuk membuat instance RecipeAPI baru
func NewRecipeAPI(baseURL string) *RecipeAPI {
    return &RecipeAPI{
        BaseURL: baseURL,
    }
}

// Fungsi untuk mencari resep berdasarkan nama makanan
func (r *RecipeAPI) GetRecipesByName(mealName string) ([]models.Recipe, error) {
    // URL untuk endpoint pencarian berdasarkan nama
    url := fmt.Sprintf("%s/search.php?s=%s", r.BaseURL, mealName)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to fetch recipes: %s", resp.Status)
    }

    // Dekode respons JSON ke dalam struct result
    var result struct {
        Meals []models.Recipe `json:"meals"`
    }
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, err
    }

    return result.Meals, nil
}
