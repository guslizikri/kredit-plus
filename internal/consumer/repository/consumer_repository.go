package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"sigmatech-kredit-plus/internal/model"
)

type ConsumerRepository struct {
	db *sqlx.DB
}
type ConsumerRepositoryIF interface {
	CreateConsumer(ctx context.Context, consumer *model.Consumer) error
	GetConsumerByNIK(ctx context.Context, nik string) (*model.Consumer, error)
}

func NewConsumerRepository(db *sqlx.DB) ConsumerRepositoryIF {
	return &ConsumerRepository{db: db}
}

func (r *ConsumerRepository) CreateConsumer(ctx context.Context, consumer *model.Consumer) (err error) {
	query := `
        INSERT INTO consumers (
            nik, full_name, legal_name, place_of_birth, date_of_birth,
            salary, photo_ktp, photo_selfie, created_at, updated_at
        ) VALUES (
            :nik, :full_name, :legal_name, :place_of_birth, :date_of_birth,
            :salary, :photo_ktp, :photo_selfie, :created_at, :updated_at
        )
    `

	_, err = r.db.NamedExecContext(ctx, query, consumer)
	if err != nil {
		return err
	}

	return nil
}

func (r *ConsumerRepository) GetConsumerByNIK(ctx context.Context, nik string) (*model.Consumer, error) {
	var consumer model.Consumer
	err := r.db.GetContext(ctx, &consumer, "SELECT * FROM consumers WHERE nik = $1", nik)
	if err != nil {
		return nil, err
	}
	return &consumer, nil
}
