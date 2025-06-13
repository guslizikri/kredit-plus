package usecase

import (
	"context"
	"sigmatech-kredit-plus/internal/limit/dto"
	"sigmatech-kredit-plus/internal/limit/repository"
	"sigmatech-kredit-plus/internal/model"
	"time"
)

type LimitUsecaseIF interface {
	CreateLimit(ctx context.Context, body *dto.CreateLimit) error
}

type LimitUsecase struct {
	repo repository.LimitRepositoryIF
}

func NewLimitUsecase(r repository.LimitRepositoryIF) *LimitUsecase {
	return &LimitUsecase{repo: r}
}

func (u *LimitUsecase) CreateLimit(ctx context.Context, body *dto.CreateLimit) error {
	var err error

	limit := model.Limit{
		UserID:      body.UserID,
		TenorMonths: body.TenorMonths,
		LimitAmount: body.LimitAmount,
		UsedAmount:  body.UsedAmount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = u.repo.CreateLimit(ctx, &limit)
	if err != nil {
		return err
	}

	return nil
}
