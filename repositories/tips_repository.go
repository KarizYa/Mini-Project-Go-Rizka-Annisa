// repositories/tips_repository.go
package repositories

import (
	"mini-project/models"
	"strings"

	"gorm.io/gorm"
)

type TipsRepository interface {
	GetAllTipsByUserID(userID uint64) ([]models.Tips, error)
	GetTipsByLeftover(userID uint64, ingredient string) ([]models.Tips, error)
	Create(tips models.Tips) error
	Update(tips models.Tips) error
	DeleteTips(userID uint64, tipID uint) error
}

type tipsRepository struct {
	db *gorm.DB
}

func NewTipsRepository(db *gorm.DB) TipsRepository {
	return &tipsRepository{
		db: db,
	}
}

// Mengambil semua tips berdasarkan userID
func (r *tipsRepository) GetAllTipsByUserID(userID uint64) ([]models.Tips, error) {
	var tips []models.Tips
	err := r.db.Where("user_id = ?", userID).Find(&tips).Error
	if err != nil {
		return nil, err
	}

	// Mengonversi Leftovers menjadi slice string setelah data diambil dari database
	for i := range tips {
		tips[i].Leftovers = strings.Join(tips[i].GetLeftoversSlice(), ",") // Update Leftovers field
	}
	return tips, nil
}

// Mengambil tips berdasarkan sisa makanan dan userID
func (r *tipsRepository) GetTipsByLeftover(userID uint64, ingredient string) ([]models.Tips, error) {
	var tips []models.Tips
	err := r.db.Where("user_id = ? AND FIND_IN_SET(?, leftovers) > 0", userID, ingredient).Find(&tips).Error
	if err != nil {
		return nil, err
	}

	// Mengonversi Leftovers menjadi slice string setelah data diambil dari database
	for i := range tips {
		tips[i].Leftovers = strings.Join(tips[i].GetLeftoversSlice(), ",") // Update Leftovers field
	}
	return tips, nil
}

// Menambahkan tips baru
func (r *tipsRepository) Create(tips models.Tips) error {
	// Mengonversi Leftovers menjadi string sebelum disimpan ke database
	tips.SetLeftoversSlice(tips.GetLeftoversSlice()) // Mengubah Leftovers menjadi string
	return r.db.Create(&tips).Error
}

// Memperbarui tips
func (r *tipsRepository) Update(tips models.Tips) error {
	// Mengonversi Leftovers menjadi string sebelum memperbarui data di database
	tips.SetLeftoversSlice(tips.GetLeftoversSlice()) // Mengubah Leftovers menjadi string
	return r.db.Save(&tips).Error
}

// Menghapus tips berdasarkan userID dan ID
func (r *tipsRepository) DeleteTips(userID uint64, tipID uint) error {
	return r.db.Where("id = ? AND user_id = ?", tipID, userID).Delete(&models.Tips{}).Error
}
