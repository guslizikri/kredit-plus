package common

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type TransactionManager interface {
	Begin(ctx context.Context) (*sqlx.Tx, error)
	Commit(ctx context.Context, tx *sqlx.Tx) error
	Rollback(ctx context.Context, tx *sqlx.Tx) error
}

type trxDBRepository struct {
	db *sqlx.DB
}

func NewTransactionManager(db *sqlx.DB) TransactionManager {
	return &trxDBRepository{db: db}
}

func (r *trxDBRepository) Begin(ctx context.Context) (*sqlx.Tx, error) {
	return r.db.BeginTxx(ctx, &sql.TxOptions{})
}

func (r *trxDBRepository) Commit(ctx context.Context, tx *sqlx.Tx) error {
	return tx.Commit()
}

func (r *trxDBRepository) Rollback(ctx context.Context, tx *sqlx.Tx) error {
	return tx.Rollback()
}
