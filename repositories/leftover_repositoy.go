package repositories

import (
	"errors"
	"mini-project/models"

	"gorm.io/gorm"
)

type LeftoverRepository interface {
    Create(leftover *models.Leftover) error
    FindAll(userID uint) ([]models.Leftover, error)
    FindByID(id uint) (models.Leftover, error)
    Update(leftover *models.Leftover) error
    Delete(id uint) error
}

type leftoverRepository struct {
    db *gorm.DB
}

func NewLeftoverRepository(db *gorm.DB) LeftoverRepository {
    return &leftoverRepository{db}
}

func (r *leftoverRepository) Create(leftover *models.Leftover) error {
    return r.db.Create(leftover).Error
}

func (r *leftoverRepository) FindAll(userID uint) ([]models.Leftover, error) {
    var leftovers []models.Leftover
    if err := r.db.Where("user_id = ?", userID).Find(&leftovers).Error; err != nil {
        return nil, err
    }
    return leftovers, nil
}

func (r *leftoverRepository) FindByID(id uint) (models.Leftover, error) {
    var leftover models.Leftover
    if err := r.db.First(&leftover, id).Error; err != nil {
        return leftover, errors.New("leftover not found")
    }
    return leftover, nil
}

func (r *leftoverRepository) Update(leftover *models.Leftover) error {
    return r.db.Save(leftover).Error
}

func (r *leftoverRepository) Delete(id uint) error {
    return r.db.Delete(&models.Leftover{}, id).Error
}
