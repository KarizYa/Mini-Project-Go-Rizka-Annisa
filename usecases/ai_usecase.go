package usecases

import (
	"mini-project/models"
	"mini-project/repositories"
)

type SuggestionUseCase struct {
	Repository *repositories.GeminiRepository
}

func NewSuggestionUseCase(repo *repositories.GeminiRepository) *SuggestionUseCase {
	return &SuggestionUseCase{
		Repository: repo,
	}
}

func (uc *SuggestionUseCase) GetSuggestion(leftover string) (*models.Suggestion, error) {
	suggestionText, err := uc.Repository.GetStorageSuggestions(leftover)
	if err != nil {
		return nil, err
	}

	return &models.Suggestion{
		Leftover:   leftover,
		Suggestion: suggestionText,
	}, nil
}
