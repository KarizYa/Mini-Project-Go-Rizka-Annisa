package main

import (
	"log"
	"mini-project/config"
	"mini-project/delivery/http"
	"mini-project/infrastructures/external"
	"mini-project/repositories"
	"mini-project/routes"
	"mini-project/usecases"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    db, err := config.InitDB()
    if err != nil {
        log.Fatal(err)
    }

    config.MigrateDatabase(db)

    userRepo := repositories.NewUserRepository(db)
    userUsecase := usecases.NewUserUsecase(userRepo)
    userHandler := http.NewUserHandler(userUsecase)

    leftoverRepo := repositories.NewLeftoverRepository(db)
    leftoverUsecase := usecases.NewLeftoverUsecase(leftoverRepo)

    leaderboardRepo := repositories.NewLeaderboardRepository(db)  
    leaderboardUsecase := usecases.NewLeaderboardUsecase(leaderboardRepo, userRepo)

    geminiRepo := repositories.NewGeminiRepository()
    suggestionUsecase := usecases.NewSuggestionUseCase(geminiRepo)

    recipeAPI := external.NewRecipeAPI("https://www.themealdb.com/api/json/v1/1")
    recipeUsecase := usecases.NewRecipeUsecase(recipeAPI)

    tipsRepo := repositories.NewTipsRepository(db)
    tipsUsecase := usecases.NewTipsUsecase(tipsRepo, userRepo)

    e := echo.New()

    routes.NewRouter(e, userHandler)

    routes.InitRoutes(e, leftoverUsecase, recipeUsecase, tipsUsecase, suggestionUsecase, leaderboardUsecase)

    log.Println("Server started on port 8000")
    if err := e.Start(":8000"); err != nil {
        log.Fatal(err)
    }
}
