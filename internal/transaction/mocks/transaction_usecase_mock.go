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

func (m *TransactionUsecaseMock) GetTransactionHistory(ctx context.Context, params dto.GetTransactionHistoryQuery) ([]*dto.GetTransactionHistoryResponse, int, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]*dto.GetTransactionHistoryResponse), args.Get(1).(int), args.Error(2)
}
