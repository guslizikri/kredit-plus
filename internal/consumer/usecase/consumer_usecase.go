package usecase

import (
	"context"
	"database/sql"
	"sigmatech-kredit-plus/internal/consumer/dto"
	"sigmatech-kredit-plus/internal/consumer/repository"
	"sigmatech-kredit-plus/internal/model"
	"sigmatech-kredit-plus/util"
	"time"
)

type ConsumerUsecaseIF interface {
	CreateConsumer(ctx context.Context, body *dto.CreateConsumer) error
	GetConsumerByNIK(ctx context.Context, nik string) (*model.Consumer, error)
}

type ConsumerUsecase struct {
	repo repository.ConsumerRepositoryIF
}

func NewConsumerUsecase(r repository.ConsumerRepositoryIF) *ConsumerUsecase {
	return &ConsumerUsecase{repo: r}
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

func (u *ConsumerUsecase) GetConsumerByNIK(ctx context.Context, nik string) (*model.Consumer, error) {
	return u.repo.GetConsumerByNIK(ctx, nik)
}
