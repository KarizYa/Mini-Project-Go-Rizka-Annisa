package repositories

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiRepository struct {
	Client *genai.Client
}

func NewGeminiRepository() *GeminiRepository {
	apiKey := os.Getenv("GEMINI_API_KEY") 
	if apiKey == "" {
		log.Fatal("Gemini API key is missing")
	}

	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini AI client: %v", err)
	}

	return &GeminiRepository{
		Client: client,
	}
}

func (repo *GeminiRepository) GetStorageSuggestions(leftover string) (string, error) {
    model := repo.Client.GenerativeModel("gemini-1.5-flash")
    resp, err := model.GenerateContent(context.Background(), genai.Text(fmt.Sprintf("How should I store %s?", leftover)))
    if err != nil {
        return "", fmt.Errorf("failed to fetch suggestions: %v", err)
    }

    for _, cand := range resp.Candidates {
        if cand.Content != nil {
            for _, part := range cand.Content.Parts {
                return fmt.Sprint(part), nil 
            }
        }
    }

    return "", fmt.Errorf("no suggestions found")
}
