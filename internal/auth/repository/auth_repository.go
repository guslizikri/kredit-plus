package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"sigmatech-kredit-plus/internal/model"
)

type AuthRepository struct {
	db *sqlx.DB
}
type AuthRepositoryIF interface {
	GetConsumerByNIK(ctx context.Context, nik string) (*model.Consumer, error)
}

func NewAuthRepository(db *sqlx.DB) AuthRepositoryIF {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetConsumerByNIK(ctx context.Context, nik string) (*model.Consumer, error) {
	var consumer model.Consumer
	err := r.db.GetContext(ctx, &consumer, "SELECT * FROM consumers WHERE nik = $1", nik)
	if err != nil {
		return nil, err
	}
	return &consumer, nil
}
