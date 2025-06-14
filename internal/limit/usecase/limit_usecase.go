package usecase

import (
	"context"
	"sigmatech-kredit-plus/internal/limit/dto"
	"sigmatech-kredit-plus/internal/limit/repository"
	"sigmatech-kredit-plus/internal/model"
	"time"
)

type LimitUsecaseIF interface {
	SetLimit(ctx context.Context, consumerId string, body *dto.SetLimit) error
}

type LimitUsecase struct {
	repo repository.LimitRepositoryIF
}

func NewLimitUsecase(r repository.LimitRepositoryIF) *LimitUsecase {
	return &LimitUsecase{repo: r}
}

func (u *LimitUsecase) SetLimit(ctx context.Context, consumerId string, body *dto.SetLimit) error {
	var err error
	exists, err := u.repo.Exists(ctx, consumerId, body.TenorMonth)
	if err != nil {
		return err
	}

	limit := model.Limit{
		ConsumerID:  consumerId,
		TenorMonth:  body.TenorMonth,
		LimitAmount: body.LimitAmount,
		UsedAmount:  0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if exists {
		return u.repo.UpdateLimit(ctx, limit.ConsumerID, limit.TenorMonth, limit.LimitAmount)
	} else {
		return u.repo.CreateLimit(ctx, &limit)
	}
}
