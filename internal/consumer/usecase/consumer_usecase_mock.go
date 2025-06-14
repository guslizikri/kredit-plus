package usecase

import (
	"context"
	"sigmatech-kredit-plus/internal/consumer/dto"
	"sigmatech-kredit-plus/internal/model"

	"github.com/stretchr/testify/mock"
)

type ConsumerUsecaseMock struct {
	mock.Mock
}

func (m *ConsumerUsecaseMock) CreateConsumer(ctx context.Context, body *dto.CreateConsumer) error {
	args := m.Called(ctx, body)
	return args.Error(0)
}

func (m *ConsumerUsecaseMock) GetConsumerByNIK(ctx context.Context, nik string) (*model.Consumer, error) {
	args := m.Called(ctx, nik)
	return args.Get(0).(*model.Consumer), args.Error(1)
}
