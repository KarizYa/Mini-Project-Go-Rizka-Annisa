// usecases/tips_usecase.go
package usecases

import (
	"mini-project/models"
	"mini-project/repositories"
)

type TipsUsecase struct {
    TipsRepo repositories.TipsRepository
}

func NewTipsUsecase(tipsRepo repositories.TipsRepository) *TipsUsecase {
    return &TipsUsecase{
        TipsRepo: tipsRepo,
    }
}

// Menampilkan semua tips berdasarkan userID
func (u *TipsUsecase) GetAllTips(userID uint64) ([]models.Tips, error) {
    return u.TipsRepo.GetAllTipsByUserID(userID)
}

// Menampilkan tips berdasarkan sisa makanan
func (u *TipsUsecase) GetTipsByLeftover(userID uint64, ingredient string) ([]models.Tips, error) {
    return u.TipsRepo.GetTipsByLeftover(userID, ingredient)
}

// Menambahkan tips baru
func (uc *TipsUsecase) CreateTips(tips models.Tips) error {
    return uc.TipsRepo.Create(tips)
}

// Memperbarui tips
func (uc *TipsUsecase) UpdateTips(tips models.Tips) error {
    return uc.TipsRepo.Update(tips)
}

// Menghapus tips
func (u *TipsUsecase) DeleteTips(userID uint64, tipID uint) error {
    return u.TipsRepo.DeleteTips(userID, tipID)
}
