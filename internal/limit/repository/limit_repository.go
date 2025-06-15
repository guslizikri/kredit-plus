package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"sigmatech-kredit-plus/internal/model"
)

type LimitRepository struct {
	db *sqlx.DB
}
type LimitRepositoryIF interface {
	Exists(ctx context.Context, consumerID string, tenor int) (bool, error)
	CreateLimit(ctx context.Context, limit *model.Limit) error
	UpdateLimit(ctx context.Context, consumerID string, tenor int, limitAmount int) error

	GetLimitWithLock(ctx context.Context, tx *sqlx.Tx, consumerID string, tenor int) (*model.Limit, error)
	UpdateUsedAmountWithTx(ctx context.Context, tx *sqlx.Tx, limitID string, usedAmount int) error
	GetLimitByConsumerID(ctx context.Context, consumerId string) ([]*model.Limit, error)
}

func NewLimitRepository(db *sqlx.DB) LimitRepositoryIF {
	return &LimitRepository{db: db}
}

func (r *LimitRepository) Exists(ctx context.Context, consumerID string, tenor int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM limits WHERE consumer_id = $1 AND tenor_month = $2)`
	err := r.db.GetContext(ctx, &exists, query, consumerID, tenor)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *LimitRepository) CreateLimit(ctx context.Context, limit *model.Limit) (err error) {
	query := `
	INSERT INTO limits (
		consumer_id, tenor_month, limit_amount, used_amount, created_at, updated_at
	) VALUES (
		:consumer_id, :tenor_month, :limit_amount, :used_amount, :created_at, :updated_at
	)
	`

	_, err = r.db.NamedExecContext(ctx, query, limit)
	if err != nil {
		return err
	}

	return nil
}

func (r *LimitRepository) UpdateLimit(ctx context.Context, consumerID string, tenor int, limitAmount int) error {
	query := `
		UPDATE limits 
		SET limit_amount = $3, updated_at = NOW()
		WHERE consumer_id = $1 AND tenor_month = $2
	`
	_, err := r.db.ExecContext(ctx, query, consumerID, tenor, limitAmount)
	return err
}

func (r *LimitRepository) GetLimitWithLock(ctx context.Context, tx *sqlx.Tx, consumerID string, tenor int) (*model.Limit, error) {
	var limit model.Limit
	query := `
		SELECT id, consumer_id, tenor_month, limit_amount, used_amount, created_at, updated_at
		FROM limits
		WHERE consumer_id = $1 AND tenor_month = $2
		FOR UPDATE
	`
	err := tx.GetContext(ctx, &limit, query, consumerID, tenor)
	return &limit, err
}

func (r *LimitRepository) UpdateUsedAmountWithTx(ctx context.Context, tx *sqlx.Tx, limitID string, usedAmount int) error {
	query := `
		UPDATE limits 
		SET used_amount = used_amount + $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err := tx.ExecContext(ctx, query, usedAmount, limitID)
	return err
}

func (r *LimitRepository) GetLimitByConsumerID(ctx context.Context, consumerId string) ([]*model.Limit, error) {
	var limits []*model.Limit
	query := `SELECT id, consumer_id, tenor_month, limit_amount, used_amount, created_at, updated_at
			  FROM limits WHERE consumer_id = $1`
	err := r.db.SelectContext(ctx, &limits, query, consumerId)
	return limits, err
}
