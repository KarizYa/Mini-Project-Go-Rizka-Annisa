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

func (u *TipsUsecase) GetAllTips() ([]models.Tips, error) {
    return u.TipsRepo.GetAllTips() 
}

func (u *TipsUsecase) GetTipsByLeftover(ingredient string) ([]models.Tips, error) {
    return u.TipsRepo.GetTipsByLeftover(ingredient)
}

func (uc *TipsUsecase) CreateTips(tips models.Tips) error {
    return uc.TipsRepo.Create(tips)
}

func (uc *TipsUsecase) UpdateTips(tips models.Tips) error {
    return uc.TipsRepo.Update(tips)
}

func (u *TipsUsecase) DeleteTips(userID uint, tipID uint) error {
    return u.TipsRepo.DeleteTips(userID, tipID)
}
