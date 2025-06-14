package usecase

import (
	"context"
	"database/sql"
	"sigmatech-kredit-plus/internal/auth/dto"
	"sigmatech-kredit-plus/internal/auth/repository"
	"sigmatech-kredit-plus/pkg"
	"sigmatech-kredit-plus/util"
)

type AuthUsecaseIF interface {
	ConsumerLogin(ctx context.Context, body *dto.ConsumerLogin) (token string, err error)
	AdminLogin(ctx context.Context, body *dto.AdminLogin) (token string, err error)
}

type AuthUsecase struct {
	repo repository.AuthRepositoryIF
}

func NewAuthUsecase(r repository.AuthRepositoryIF) *AuthUsecase {
	return &AuthUsecase{repo: r}
}

func (u *AuthUsecase) ConsumerLogin(ctx context.Context, body *dto.ConsumerLogin) (token string, err error) {

	consumer, err := u.repo.GetConsumerByNIK(ctx, body.NIK)
	if err != nil && err != sql.ErrNoRows {
		return "", util.InternalServerError(err.Error())
	}
	if err == sql.ErrNoRows {
		return "", util.NotFound("nik doesnt exist")
	}
	if consumer.FullName != body.FullName {
		return "", util.NotFound("name invalid")
	}

	jwtt := pkg.NewToken(consumer.ID, "", "consumer")
	tokens, err := jwtt.Generate()
	if err != nil {
		return "", err
	}

	return tokens, nil
}

func (u *AuthUsecase) AdminLogin(ctx context.Context, body *dto.AdminLogin) (token string, err error) {
	// for now admin login is still hardcode
	jwtt := pkg.NewToken("", body.Username, "admin")
	tokens, err := jwtt.Generate()
	if err != nil {
		return "", err
	}

	return tokens, nil
}
