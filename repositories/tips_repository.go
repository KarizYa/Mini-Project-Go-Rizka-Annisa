package repositories

import (
	"mini-project/models"
	"gorm.io/gorm"
)

type TipsRepository interface {
    GetAllTips() ([]models.Tips, error) 
    GetTipsByLeftover(ingredient string) ([]models.Tips, error) 
    Create(tips models.Tips) error
    Update(tips models.Tips) error
    DeleteTips(userID uint, tipID uint) error
}

type tipsRepository struct {
    DB *gorm.DB
}

func NewTipsRepository(DB *gorm.DB) TipsRepository {
    return &tipsRepository{DB}
}

func (r *tipsRepository) GetAllTips() ([]models.Tips, error) {
    var tips []models.Tips
    if err := r.DB.Find(&tips).Error; err != nil {
        return nil, err
    }
    return tips, nil
}

func (r *tipsRepository) GetTipsByLeftover(leftover string) ([]models.Tips, error) {
    var tips []models.Tips
    if err := r.DB.Where("leftovers LIKE ?", "%"+leftover+"%").Find(&tips).Error; err != nil {
        return nil, err
    }
    return tips, nil
}


func (r *tipsRepository) Create(tips models.Tips) error {
    return r.DB.Create(&tips).Error
}

func (r *tipsRepository) Update(tips models.Tips) error {
    return r.DB.Save(&tips).Error
}

func (r *tipsRepository) DeleteTips(userID uint, tipID uint) error {
    return r.DB.Where("user_id = ? AND id = ?", userID, tipID).Delete(&models.Tips{}).Error
}
