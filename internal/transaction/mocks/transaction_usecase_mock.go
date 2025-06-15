package mocks

import (
	"context"
	"sigmatech-kredit-plus/internal/transaction/dto"

	"github.com/stretchr/testify/mock"
)

type TransactionUsecaseMock struct {
	mock.Mock
}

func (m *TransactionUsecaseMock) CreateTransaction(ctx context.Context, body *dto.CreateTransaction, consumerId string) (string, error) {
	args := m.Called(ctx, body, consumerId)
	return args.Get(0).(string), args.Error(1)
}
