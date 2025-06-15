package usecase

import (
	"context"
	"database/sql"
	"sigmatech-kredit-plus/internal/consumer/dto"
	"sigmatech-kredit-plus/internal/consumer/repository"
	limit_repository "sigmatech-kredit-plus/internal/limit/repository"
	"sigmatech-kredit-plus/internal/model"
	"sigmatech-kredit-plus/util"
	"time"
)

type ConsumerUsecaseIF interface {
	CreateConsumer(ctx context.Context, body *dto.CreateConsumer) error
	GetConsumerDetail(ctx context.Context, id string) (*dto.GetConsumerDetailResponse, error)
}

type ConsumerUsecase struct {
	repo      repository.ConsumerRepositoryIF
	limitRepo limit_repository.LimitRepositoryIF
}

func NewConsumerUsecase(repo repository.ConsumerRepositoryIF, limitRepo limit_repository.LimitRepositoryIF) *ConsumerUsecase {
	return &ConsumerUsecase{
		repo:      repo,
		limitRepo: limitRepo,
	}
}

func (u *ConsumerUsecase) CreateConsumer(ctx context.Context, body *dto.CreateConsumer) error {
	var err error

	_, err = u.repo.GetConsumerByNIK(ctx, body.NIK)
	if err != nil && err != sql.ErrNoRows {
		return util.InternalServerError(err.Error())
	}
	if err == nil {
		return util.Conflict("nik already registered")
	}

	consumer := model.Consumer{
		NIK:         body.NIK,
		FullName:    body.FullName,
		LegalName:   body.LegalName,
		BirthPlace:  body.BirthPlace,
		BirthDate:   body.BirthDate,
		Salary:      body.Salary,
		PhotoKTP:    body.PhotoKTP,
		PhotoSelfie: body.PhotoSelfie,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = u.repo.CreateConsumer(ctx, &consumer)
	if err != nil {
		return err
	}

	return nil
}

func (u *ConsumerUsecase) GetConsumerDetail(ctx context.Context, id string) (*dto.GetConsumerDetailResponse, error) {
	consumer, err := u.repo.GetConsumerByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, util.NotFound("consumer not found")
		}
		return nil, util.InternalServerError(err.Error())
	}

	limits, err := u.limitRepo.GetLimitByConsumerID(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, l := range limits {
		consumer.Limits = append(consumer.Limits, dto.LimitEmbedded{
			TenorMonth:      l.TenorMonth,
			LimitAmount:     l.LimitAmount,
			UsedAmount:      l.UsedAmount,
			AvailableAmount: l.LimitAmount - l.UsedAmount,
		})
	}

	return consumer, nil
}
