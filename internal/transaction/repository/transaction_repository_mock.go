package repository

import (
	"context"
	"sigmatech-kredit-plus/internal/model"

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
