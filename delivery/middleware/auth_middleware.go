// package middleware
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

// init fungsi ini akan secara otomatis memuat variabel lingkungan dari file .env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// JWTAuthMiddleware memastikan bahwa hanya pengguna yang telah login dapat mengakses endpoint tertentu.
// Middleware ini akan memverifikasi token JWT dan mengeluarkan error Unauthorized jika token tidak valid.
func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Mengambil header Authorization
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Authorization header missing")
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		// Menghapus prefix "Bearer " dari token
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Mendapatkan secret key dari variabel lingkungan
		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			log.Println("JWT secret key not found in environment variables")
			return echo.NewHTTPError(http.StatusInternalServerError, "JWT_SECRET not found in environment variables")
		}

		// Parsing dan memvalidasi token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Memastikan token menggunakan metode signing yang benar
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("Unexpected signing method: %v", token.Method)
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(secretKey), nil
		})

		// Log token jika parsing berhasil
		if token != nil {
			log.Printf("Parsed Token: %v", token)
		}

		// Jika ada error atau token tidak valid, kembalikan error Unauthorized
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}
		if !token.Valid {
			log.Println("Token is not valid")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		// Mengekstrak user_id dari klaim token dan menyimpannya dalam konteks Echo
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Mendapatkan user_id dari klaim dan menyimpannya ke dalam konteks
			userID := uint(claims["user_id"].(float64)) // Melakukan cast user_id ke tipe uint
			c.Set("userID", userID) // userID ini dapat digunakan di handler yang memerlukan ID pengguna

			// Log user ID yang berhasil diambil
			log.Printf("User ID extracted from token: %v", userID)

			// Opsional: log klaim token lainnya (seperti email)
			if email, ok := claims["email"].(string); ok {
				log.Printf("Email extracted from token: %v", email)
			} else {
				log.Println("Email claim not found in token")
			}
		} else {
			log.Println("Invalid token claims")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}

		// Melanjutkan ke handler berikutnya
		return next(c)
	}
}
