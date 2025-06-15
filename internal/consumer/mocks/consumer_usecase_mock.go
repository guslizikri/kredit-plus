package mocks

import (
	"context"
	"sigmatech-kredit-plus/internal/consumer/dto"

	"github.com/stretchr/testify/mock"
)

type ConsumerUsecaseMock struct {
	mock.Mock
}

func (m *ConsumerUsecaseMock) CreateConsumer(ctx context.Context, body *dto.CreateConsumer) error {
	args := m.Called(ctx, body)
	return args.Error(0)
}

func (m *ConsumerUsecaseMock) GetConsumerDetail(ctx context.Context, id string) (*dto.GetConsumerDetailResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*dto.GetConsumerDetailResponse), args.Error(1)
}
