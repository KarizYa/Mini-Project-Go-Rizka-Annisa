package repositories

import (
	"mini-project/models"
	"gorm.io/gorm"
)

type TipsRepository interface {
    GetAllTips() ([]models.Tips, error) // Ambil semua tips tanpa filter userID
    GetTipsByLeftover(ingredient string) ([]models.Tips, error) // Ambil tips berdasarkan bahan makanan tanpa filter userID
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

// Fungsi untuk mendapatkan semua tips tanpa filter berdasarkan userID
func (r *tipsRepository) GetAllTips() ([]models.Tips, error) {
    var tips []models.Tips
    if err := r.DB.Find(&tips).Error; err != nil {
        return nil, err
    }
    return tips, nil
}

// Fungsi untuk mendapatkan tips berdasarkan sisa makanan tanpa filter userID
func (r *tipsRepository) GetTipsByLeftover(ingredient string) ([]models.Tips, error) {
    var tips []models.Tips
    if err := r.DB.Where("leftovers LIKE ?", "%"+ingredient+"%").Find(&tips).Error; err != nil {
        return nil, err
    }
    return tips, nil
}

// Fungsi untuk menambahkan tips baru
func (r *tipsRepository) Create(tips models.Tips) error {
    return r.DB.Create(&tips).Error
}

// Fungsi untuk memperbarui tips
func (r *tipsRepository) Update(tips models.Tips) error {
    return r.DB.Save(&tips).Error
}

// Fungsi untuk menghapus tips
func (r *tipsRepository) DeleteTips(userID uint, tipID uint) error {
    return r.DB.Where("user_id = ? AND id = ?", userID, tipID).Delete(&models.Tips{}).Error
}
