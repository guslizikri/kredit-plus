package dto

import "time"

type CreateUser struct {
	NIK          string    `json:"nik" db:"nik" validate:"required"`
	FullName     string    `json:"full_name" db:"full_name" validate:"required"`
	LegalName    string    `json:"legal_name" db:"legal_name" validate:"required"`
	PlaceOfBirth string    `json:"place_of_birth" db:"place_of_birth" validate:"-"`
	DateOfBirth  time.Time `json:"date_of_birth" db:"date_of_birth" time_format:"2006-01-02" validate:"-"`
	Salary       float64   `json:"salary" db:"salary" validate:"required"`
	PhotoKTP     string    `json:"photo_ktp" db:"photo_ktp" validate:"-"`
	PhotoSelfie  string    `json:"photo_selfie" db:"photo_selfie" validate:"-"`
}
