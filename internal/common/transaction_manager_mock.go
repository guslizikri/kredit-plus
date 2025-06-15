package common

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type TransactionManagerMock struct {
	mock.Mock
}

func (m *TransactionManagerMock) Begin(ctx context.Context) (*sqlx.Tx, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sqlx.Tx), args.Error(1)
}
func (m *TransactionManagerMock) Commit(ctx context.Context, tx *sqlx.Tx) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}
func (m *TransactionManagerMock) Rollback(ctx context.Context, tx *sqlx.Tx) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}
