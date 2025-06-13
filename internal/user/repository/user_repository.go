package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"sigmatech-kredit-plus/internal/model"
)

type UserRepository struct {
	db *sqlx.DB
}
type UserRepositoryIF interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByNIK(ctx context.Context, nik string) (*model.User, error)
}

func NewUserRepository(db *sqlx.DB) UserRepositoryIF {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (err error) {
	query := `
        INSERT INTO users (
            nik, full_name, legal_name, place_of_birth, date_of_birth,
            salary, photo_ktp, photo_selfie, created_at, updated_at
        ) VALUES (
            :nik, :full_name, :legal_name, :place_of_birth, :date_of_birth,
            :salary, :photo_ktp, :photo_selfie, :created_at, :updated_at
        )
    `

	_, err = r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByNIK(ctx context.Context, nik string) (*model.User, error) {
	var user model.User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE nik = $1", nik)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
