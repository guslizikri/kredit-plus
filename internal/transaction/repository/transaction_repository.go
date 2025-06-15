package repository

import (
	"context"
	"sigmatech-kredit-plus/internal/model"
	"sigmatech-kredit-plus/internal/transaction/dto"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryIF interface {
	CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx *model.Transaction) error
	FetchTransactionByConsumer(ctx context.Context, consumerID string, limit, offset int) ([]*dto.GetTransactionHistoryResponse, error)
	CountTransactionByConsumer(ctx context.Context, consumerID string) (int, error)
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

func (r *TransactionRepository) FetchTransactionByConsumer(ctx context.Context, consumerID string, limit, offset int) ([]*dto.GetTransactionHistoryResponse, error) {
	var result []*dto.GetTransactionHistoryResponse
	query := `
		SELECT 
			t.contract_number, t.otr_price, t.installment, t.interest, t.asset_name, t.created_at,
			l.tenor_month
		FROM transactions t
		JOIN limits l ON l.id = t.limit_id
		WHERE t.consumer_id = $1
		ORDER BY t.created_at DESC
		LIMIT $2 OFFSET $3
	`
	err := r.db.SelectContext(ctx, &result, query, consumerID, limit, offset)
	return result, err
}

func (r *TransactionRepository) CountTransactionByConsumer(ctx context.Context, consumerID string) (int, error) {
	var total int
	query := `SELECT COUNT(1) FROM transactions WHERE consumer_id = $1`
	err := r.db.GetContext(ctx, &total, query, consumerID)
	return total, err
}
