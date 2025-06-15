package mocks

import (
	"context"
	"sigmatech-kredit-plus/internal/model"
	"sigmatech-kredit-plus/internal/transaction/dto"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (m *RepoMock) CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx *model.Transaction) error {
	args := m.Called(ctx, tx, trx)
	return args.Error(0)
}
func (m *RepoMock) FetchTransactionByConsumer(ctx context.Context, consumerID string, limit, offset int) ([]*dto.GetTransactionHistoryResponse, error) {
	args := m.Called(ctx, consumerID, limit, offset)
	return args.Get(0).([]*dto.GetTransactionHistoryResponse), args.Error(1)
}
func (m *RepoMock) CountTransactionByConsumer(ctx context.Context, consumerID string) (int, error) {
	args := m.Called(ctx, consumerID)
	return args.Get(0).(int), args.Error(1)
}
