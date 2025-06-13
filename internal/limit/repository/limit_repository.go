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
	CreateLimit(ctx context.Context, limit *model.Limit) error
}

func NewLimitRepository(db *sqlx.DB) LimitRepositoryIF {
	return &LimitRepository{db: db}
}

func (r *LimitRepository) CreateLimit(ctx context.Context, limit *model.Limit) (err error) {
	query := `
	INSERT INTO limits (
		id, limit_id, tenor_months, limit_amount, used_amount, created_at, updated_at
	) VALUES (
		gen_random_uuid(), :limit_id, :tenor_months, :limit_amount, :used_amount, :created_at, :updated_at
	)
	`

	_, err = r.db.NamedExecContext(ctx, query, limit)
	if err != nil {
		return err
	}

	return nil
}
