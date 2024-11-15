package models

// Struktur data untuk resep yang disesuaikan dengan API TheMealDB
type Recipe struct {
	IDMeal       string         `json:"idMeal"`
	Title        string         `json:"strMeal"`
	Category     string         `json:"strCategory"`
	Area         string         `json:"strArea"`
	Instructions string         `json:"strInstructions"`
	MealThumb    string         `json:"strMealThumb"`
	Tags         string         `json:"strTags"`
	Youtube      string         `json:"strYoutube"`
	Ingredients  []string       `json:"strIngredients" gorm:"type:json"` 
}
