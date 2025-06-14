package model

import "time"

type Consumer struct {
	ID          string    `json:"id" db:"id"`
	NIK         string    `json:"nik" db:"nik"`
	FullName    string    `json:"full_name" db:"full_name"`
	LegalName   string    `json:"legal_name" db:"legal_name"`
	BirthPlace  string    `json:"birth_place" db:"birth_place"`
	BirthDate   time.Time `json:"birth_date" db:"birth_date"`
	Salary      int       `json:"salary" db:"salary"`
	PhotoKTP    string    `json:"photo_ktp" db:"photo_ktp"`
	PhotoSelfie string    `json:"photo_selfie" db:"photo_selfie"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
