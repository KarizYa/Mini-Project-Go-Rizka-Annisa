package http

import (
	"mini-project/models"
	"mini-project/usecases"
	"mini-project/helper" 
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
    userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
    return &UserHandler{userUsecase}
}

func (h *UserHandler) Register(c echo.Context) error {
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, helper.WrapResponse("Invalid input", 400, "error", nil))
    }

    if user.Email == "" || user.Password == "" {
        return c.JSON(http.StatusBadRequest, helper.WrapResponse("Email and password are required", 400, "error", nil))
    }

    if err := h.userUsecase.Register(user); err != nil {
        return c.JSON(http.StatusInternalServerError, helper.WrapResponse(err.Error(), 500, "error", nil))
    }

    return c.JSON(http.StatusCreated, helper.WrapResponse("User registered successfully", 201, "success", map[string]interface{}{
        "email": user.Email,
        "password": "",
    }))
}

func (h *UserHandler) Login(c echo.Context) error {
    var credentials models.Credentials
    if err := c.Bind(&credentials); err != nil {
        return c.JSON(http.StatusBadRequest, helper.WrapResponse("Invalid input", 400, "error", nil))
    }

    if credentials.Email == "" || credentials.Password == "" {
        return c.JSON(http.StatusBadRequest, helper.WrapResponse("Email and password are required", 400, "error", nil))
    }

    token, err := h.userUsecase.Login(credentials)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.WrapResponse(err.Error(), 401, "error", nil))
    }

    return c.JSON(http.StatusOK, helper.WrapResponse("Login successful", 200, "success", map[string]interface{}{
        "email": credentials.Email,
        "password": "",  
        "token": token,
    }))
}
