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

// Menampilkan semua tips tanpa filter berdasarkan userID
func (u *TipsUsecase) GetAllTips() ([]models.Tips, error) {
    // Ambil semua tips tanpa memfilter berdasarkan userID
    return u.TipsRepo.GetAllTips() 
}

// Menampilkan tips berdasarkan sisa makanan tanpa filter userID
func (u *TipsUsecase) GetTipsByLeftover(ingredient string) ([]models.Tips, error) {
    // Ambil tips berdasarkan sisa makanan tanpa memfilter berdasarkan userID
    return u.TipsRepo.GetTipsByLeftover(ingredient)
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
func (u *TipsUsecase) DeleteTips(userID uint, tipID uint) error {
    return u.TipsRepo.DeleteTips(userID, tipID)
}
