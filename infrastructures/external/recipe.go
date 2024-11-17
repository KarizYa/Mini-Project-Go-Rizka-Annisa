package external

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"mini-project/models"
	"net/http"
)

type RecipeAPI struct {
	BaseURL string
}

func NewRecipeAPI(baseURL string) *RecipeAPI {
	return &RecipeAPI{
		BaseURL: baseURL,
	}
}

func (r *RecipeAPI) GetRecipesByName(mealName string) ([]models.Recipe, error) {
	url := fmt.Sprintf("%s/search.php?s=%s", r.BaseURL, mealName)

	// Custom HTTP client with TLS skipping
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	// Making the GET request
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch recipes: %s", resp.Status)
	}

	var result struct {
		Meals []models.Recipe `json:"meals"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Meals, nil
}
