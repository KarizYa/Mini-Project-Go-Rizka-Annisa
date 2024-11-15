package main

import (
	"log"
	"mini-project/config"
	"mini-project/delivery/http"
	"mini-project/infrastructures/external"
	"mini-project/repositories"
	"mini-project/routes"
	"mini-project/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
    db, err := config.InitDB()
    if err != nil {
        log.Fatal(err)
    }

    config.MigrateDatabase(db)

    // Inisialisasi repository dan usecase untuk user
    userRepo := repositories.NewUserRepository(db)
    userUsecase := usecases.NewUserUsecase(userRepo)
    userHandler := http.NewUserHandler(userUsecase)

    // Inisialisasi repository dan usecase untuk leftover
    leftoverRepo := repositories.NewLeftoverRepository(db)
    leftoverUsecase := usecases.NewLeftoverUsecase(leftoverRepo)

    // Inisialisasi RecipeAPI untuk menggunakan API TheMealDB
    recipeAPI := external.NewRecipeAPI("https://www.themealdb.com/api/json/v1/1") // Base URL API
    recipeUsecase := usecases.NewRecipeUsecase(recipeAPI)

    // Inisialisasi repository dan usecase untuk tips
    tipsRepo := repositories.NewTipsRepository(db)
    tipsUsecase := usecases.NewTipsUsecase(tipsRepo) // Perbaikan: Gunakan tipsUsecase

    // Membuat echo instance
    e := echo.New()

    // Setup routes untuk user
    routes.NewRouter(e, userHandler)

    // Menambahkan leftoverUsecase, recipeUsecase, dan tipsUsecase ke routes
    routes.InitRoutes(e, leftoverUsecase, recipeUsecase, tipsUsecase) // Kirimkan tipsUsecase ke InitRoutes

    log.Println("Server started on port 8000")
    if err := e.Start(":8000"); err != nil {
        log.Fatal(err)
    }
}
