package repository

import (
	"github.com/jmoiron/sqlx"

	"sigmatech-kredit-plus/internal/user/dto"
)

type PgUserRepository struct {
	DB *sqlx.DB
}
type UserRepository interface {
	CreateUser(user *dto.User) error
	GetUserByID(id string) (*dto.User, error)
}

func NewPgUserRepository(db *sqlx.DB) UserRepository {
	return &PgUserRepository{DB: db}
}

func (r *PgUserRepository) CreateUser(user *dto.User) error {
	query := `
        INSERT INTO users (
            id, nik, full_name, legal_name, place_of_birth, date_of_birth,
            salary, photo_ktp, photo_selfie, created_at, updated_at
        ) VALUES (
            gen_random_uuid(), :nik, :full_name, :legal_name, :place_of_birth, :date_of_birth,
            :salary, :photo_ktp, :photo_selfie, NOW(), NOW()
        ) RETURNING id
    `
	stmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return err
	}
	return stmt.Get(&user.ID, user)
}

func (r *PgUserRepository) GetUserByID(id string) (*dto.User, error) {
	var user dto.User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	return &user, err
}
