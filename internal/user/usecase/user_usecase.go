package usecase

import (
	"context"
	"database/sql"
	"sigmatech-kredit-plus/internal/model"
	"sigmatech-kredit-plus/internal/user/dto"
	"sigmatech-kredit-plus/internal/user/repository"
	"sigmatech-kredit-plus/util"
	"time"
)

type UserUsecase struct {
	repo repository.UserRepositoryIF
}

func NewUserUsecase(r repository.UserRepositoryIF) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) CreateUser(ctx context.Context, body *dto.CreateUser) error {
	var err error

	_, err = u.repo.GetUserByNIK(ctx, body.NIK)
	if err != nil && err != sql.ErrNoRows {
		return util.InternalServerError(err.Error())
	}
	if err == nil {
		return util.Conflict("nik already registered")
	}

	user := model.User{
		NIK:          body.NIK,
		FullName:     body.FullName,
		LegalName:    body.LegalName,
		PlaceOfBirth: body.PlaceOfBirth,
		DateOfBirth:  body.DateOfBirth,
		Salary:       body.Salary,
		PhotoKTP:     body.PhotoKTP,
		PhotoSelfie:  body.PhotoSelfie,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = u.repo.CreateUser(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) GetUserByNIK(ctx context.Context, nik string) (*model.User, error) {
	return u.repo.GetUserByNIK(ctx, nik)
}
