package mocks

import (
	"context"
	"sigmatech-kredit-plus/internal/consumer/dto"
	"sigmatech-kredit-plus/internal/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) CreateConsumer(ctx context.Context, consumer *model.Consumer) error {
	args := r.Mock.Called(ctx, consumer)
	return args.Error(0)
}

func (r *RepoMock) GetConsumerByNIK(ctx context.Context, nik string) (*model.Consumer, error) {
	args := r.Mock.Called(ctx, nik)
	return args.Get(0).(*model.Consumer), args.Error(1)
}

func (r *RepoMock) GetConsumerByID(ctx context.Context, nik string) (*dto.GetConsumerDetailResponse, error) {
	args := r.Mock.Called(ctx, nik)
	return args.Get(0).(*dto.GetConsumerDetailResponse), args.Error(1)
}
