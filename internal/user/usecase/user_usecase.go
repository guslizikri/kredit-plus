package usecase

import (
	"sigmatech-kredit-plus/internal/user/dto"
	"sigmatech-kredit-plus/internal/user/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) CreateUser(user *dto.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserUsecase) GetUserByID(id string) (*dto.User, error) {
	return u.repo.GetUserByID(id)
}
