package routes

import (
	"mini-project/delivery/http"
	"mini-project/delivery/middleware"
	"mini-project/usecases"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, userHandler *http.UserHandler) {
	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)
}

func InitRoutes(e *echo.Echo, leftoverUsecase usecases.LeftoverUsecase, recipeUsecase *usecases.RecipeUsecase, tipsUsecase *usecases.TipsUsecase) {
	// Setup Leftover Routes
	leftoverHandler := http.NewLeftoverHandler(leftoverUsecase)
	leftoverGroup := e.Group("/leftovers", middleware.JWTAuthMiddleware)

	leftoverGroup.Use(middleware.JWTAuthMiddleware)

	leftoverGroup.POST("", leftoverHandler.CreateLeftover)
	leftoverGroup.GET("", leftoverHandler.GetAllLeftovers)
	leftoverGroup.GET("/:id", leftoverHandler.GetLeftoverByID)
	leftoverGroup.PUT("/:id", leftoverHandler.UpdateLeftover)
	leftoverGroup.DELETE("/:id", leftoverHandler.DeleteLeftover)

	// Setup Recipe Routes
	recipeHandler := http.NewRecipeHandler(recipeUsecase)
	recipeGroup := e.Group("/recipes")

	// Menambahkan middleware otentikasi untuk grup resep
	recipeGroup.Use(middleware.JWTAuthMiddleware)  // Pastikan grup ini dilindungi oleh JWT

	// Endpoint untuk pencarian resep
	recipeGroup.GET("/search", recipeHandler.SearchRecipesHandler)

	// Setup Tips Routes (Hanya pengguna terautentikasi yang dapat mengakses)
	tipsHandler := http.NewTipsHandler(tipsUsecase)
	tipsGroup := e.Group("/tips", middleware.JWTAuthMiddleware)  // Menambahkan middleware JWTAuthMiddleware pada grup tips

	// Endpoint untuk mengelola tips
	tipsGroup.GET("", tipsHandler.GetAllTips)
	tipsGroup.GET("/search", tipsHandler.GetTipsByLeftover)
	tipsGroup.POST("", tipsHandler.CreateTips)
	tipsGroup.PUT("/:id", tipsHandler.UpdateTips)
	tipsGroup.DELETE("/:id", tipsHandler.DeleteTips)
}
