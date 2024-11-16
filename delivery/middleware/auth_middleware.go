package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Authorization header missing")
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			log.Println("JWT secret key not found in environment variables")
			return echo.NewHTTPError(http.StatusInternalServerError, "JWT_SECRET not found in environment variables")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("Unexpected signing method: %v", token.Method)
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(secretKey), nil
		})

		if token != nil {
			log.Printf("Parsed Token: %v", token)
		}

		if err != nil {
			log.Printf("Error parsing token: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}
		if !token.Valid {
			log.Println("Token is not valid")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := uint(claims["user_id"].(float64)) 
			c.Set("userID", userID) 

			log.Printf("User ID extracted from token: %v", userID)

			if email, ok := claims["email"].(string); ok {
				log.Printf("Email extracted from token: %v", email)
			} else {
				log.Println("Email claim not found in token")
			}
		} else {
			log.Println("Invalid token claims")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}
		return next(c)
	}
}
