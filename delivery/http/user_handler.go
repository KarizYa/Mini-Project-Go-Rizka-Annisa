package http

import (
	"mini-project/models"
	"mini-project/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
    userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
    return &UserHandler{userUsecase}
}

// Register endpoint untuk registrasi user baru
func (h *UserHandler) Register(c echo.Context) error {
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    if user.Email == "" || user.Password == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email and password are required"})
    }

    // Proses registrasi
    if err := h.userUsecase.Register(user); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

// Login endpoint untuk autentikasi dan menghasilkan token JWT
func (h *UserHandler) Login(c echo.Context) error {
    var credentials models.Credentials
    if err := c.Bind(&credentials); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    if credentials.Email == "" || credentials.Password == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email and password are required"})
    }

    // Proses login dan mendapatkan token
    token, err := h.userUsecase.Login(credentials)
    if err != nil {
        // Error handling jika email atau password tidak cocok
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]string{"token": token})
}
