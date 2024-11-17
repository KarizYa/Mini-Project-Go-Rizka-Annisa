package usecases

import (
	"mini-project/models"
	"mini-project/repositories"
	"strings"
)

type TipsUsecase struct {
    TipsRepo repositories.TipsRepository
    UserRepo repositories.UserRepository
}

func NewTipsUsecase(tipsRepo repositories.TipsRepository, userRepo repositories.UserRepository) *TipsUsecase {
    return &TipsUsecase{
        TipsRepo: tipsRepo,
        UserRepo: userRepo,
    }
}

func (u *TipsUsecase) GetAllTips() ([]models.Tips, error) {
    return u.TipsRepo.GetAllTips() 
}

func (u *TipsUsecase) GetTipsByLeftover(leftover string) ([]models.Tips, error) {
    allTips, err := u.TipsRepo.GetAllTips() 
    if err != nil {
        return nil, err
    }

    var filteredTips []models.Tips
    for _, tip := range allTips {
        leftoversList := strings.Split(tip.Leftovers, ",")
        for _, item := range leftoversList {
            if strings.TrimSpace(item) == leftover {
                filteredTips = append(filteredTips, tip)
                break
            }
        }
    }

    return filteredTips, nil
}

func (uc *TipsUsecase) CreateTips(tips models.Tips) error {
    err := uc.TipsRepo.Create(tips)
    if err != nil {
        return err
    }

    user, err := uc.UserRepo.GetByID(tips.UserID) 
    if err != nil {
        return err
    }

    user.Points += 10 
    if err := uc.UserRepo.Update(user); err != nil {
        return err
    }

    return nil
}

func (uc *TipsUsecase) UpdateTips(tips models.Tips) error {
    return uc.TipsRepo.Update(tips)
}

func (u *TipsUsecase) DeleteTips(userID uint, tipID uint) error {
    return u.TipsRepo.DeleteTips(userID, tipID)
}
