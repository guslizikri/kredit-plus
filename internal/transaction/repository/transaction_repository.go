package repository

import (
	"context"
	"sigmatech-kredit-plus/internal/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryIF interface {
	CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx *model.Transaction) error
}

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepositoryIF {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx *model.Transaction) error {
	query := `
		INSERT INTO transactions (id, consumer_id, limit_id, contract_number, otr_price, admin_fee, installment, interest, asset_name, source_channel, created_at)
		VALUES (:id, :consumer_id, :limit_id, :contract_number, :otr_price, :admin_fee, :installment, :interest, :asset_name, :source_channel, :created_at)
	`
	_, err := tx.NamedExecContext(ctx, query, trx)
	return err
}
