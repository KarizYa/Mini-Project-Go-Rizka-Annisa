package usecases

import (
	"errors"
	"log"
	"mini-project/models"
	"mini-project/repositories"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
    Register(user models.User) error
    Login(credentials models.Credentials) (string, error) 
}

type userUsecase struct {
    userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
    return &userUsecase{userRepo}
}

// Register a new user with hashed password
func (u *userUsecase) Register(user models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password for email %s", user.Email) // Debugging log
        return errors.New("failed to hash password")
    }
    user.Password = string(hashedPassword)

    // Save user to repository
    if err := u.userRepo.Register(user); err != nil {
        log.Printf("Error saving user to the database: %v", err) // Debugging log
        return err
    }

    log.Printf("Success: User registered with email %s", user.Email) // Debugging log
    return nil
}

// Login with email and password and return JWT token
func (u *userUsecase) Login(credentials models.Credentials) (string, error) {
    // Find user by email
    user, err := u.userRepo.GetByEmail(credentials.Email)
    if err != nil {
        log.Printf("Error: User not found with email %s", credentials.Email) // Debugging log
        return "", errors.New("user not found") // If user is not found
    }

    // Verify password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
    if err != nil {
        log.Printf("Error: Password mismatch for email %s", credentials.Email) // Debugging log
        return "", errors.New("incorrect password") // If password doesn't match
    }

    // Log success for password verification
    log.Printf("Success: Password matched for email %s", credentials.Email)

    // Generate JWT token with user ID if login is successful
    token, err := generateJWT(user.ID, user.Email)
    if err != nil {
        log.Printf("Error: Failed to generate token for email %s", credentials.Email) 
        return "", errors.New("failed to generate token")
    }

    log.Printf("Success: Token generated for email %s", credentials.Email) 

    return token, nil
}

// generateJWT creates a JWT token for the given user ID and email
func generateJWT(userID uint, email string) (string, error) {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file") // Log error if .env file cannot be loaded
        return "", err
    }

    // Get secret key from environment variables
    secretKey := os.Getenv("JWT_SECRET_KEY")
    if secretKey == "" {
        return "", errors.New("JWT_SECRET_KEY is not set in .env")
    }

    // Prepare JWT claims
    claims := jwt.MapClaims{
        "user_id": userID,  // Add user_id to claims
        "email":   email,
        "exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expiration set to 72 hours
    }

    // Create new JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(secretKey))
    if err != nil {
        log.Printf("Error signing JWT token for email %s", email)
        return "", err
    }

    return signedToken, nil
}
