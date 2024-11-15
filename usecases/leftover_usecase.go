package usecases

import (
	"mini-project/models"
	"mini-project/repositories"
)

type LeftoverUsecase interface {
    CreateLeftover(leftover *models.Leftover) error
    GetAllLeftovers(userID uint) ([]models.Leftover, error)
    GetLeftoverByID(id uint) (models.Leftover, error)
    UpdateLeftover(leftover *models.Leftover) error
    DeleteLeftover(id uint) error
}

type leftoverUsecase struct {
    repository repositories.LeftoverRepository
}

func NewLeftoverUsecase(repository repositories.LeftoverRepository) LeftoverUsecase {
    return &leftoverUsecase{repository}
}

func (u *leftoverUsecase) CreateLeftover(leftover *models.Leftover) error {
    return u.repository.Create(leftover)
}

func (u *leftoverUsecase) GetAllLeftovers(userID uint) ([]models.Leftover, error) {
    return u.repository.FindAll(userID)
}

func (u *leftoverUsecase) GetLeftoverByID(id uint) (models.Leftover, error) {
    return u.repository.FindByID(id)
}

func (u *leftoverUsecase) UpdateLeftover(leftover *models.Leftover) error {
    return u.repository.Update(leftover)
}

func (u *leftoverUsecase) DeleteLeftover(id uint) error {
    return u.repository.Delete(id)
}
