package repositories

import (
    "errors"
    "log"
    "mini-project/models"

    "gorm.io/gorm"
)

type UserRepository interface {
    Register(user models.User) error
    GetByEmail(email string) (models.User, error)
    GetByID(userID uint) (models.User, error) 
    Update(user models.User) error          
}


type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) Register(user models.User) error {
    result := r.db.Create(&user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *userRepository) GetByEmail(email string) (models.User, error) {
    var user models.User
    result := r.db.Where("email = ? AND (deleted_at IS NULL OR deleted_at = '0000-00-00 00:00:00.000')", email).First(&user)
    
    log.Printf("Querying for user with email: %s", email)

    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            log.Printf("User not found with email: %s", email)
            return user, errors.New("user not found")
        }
        log.Printf("Error retrieving user with email: %s, Error: %v", email, result.Error)
        return user, result.Error
    }

    log.Printf("User found: %+v", user)
    return user, nil
}

func (r *userRepository) Update(user models.User) error {
    return r.db.Save(&user).Error
}

func (r *userRepository) GetByID(userID uint) (models.User, error) {
    var user models.User
    result := r.db.First(&user, userID)
    if result.Error != nil {
        return user, result.Error
    }
    return user, nil
}
