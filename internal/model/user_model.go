package model

import "time"

type Consumer struct {
	ID           string    `json:"id" db:"id"`
	NIK          string    `json:"nik" db:"nik"`
	FullName     string    `json:"full_name" db:"full_name"`
	LegalName    string    `json:"legal_name" db:"legal_name"`
	PlaceOfBirth string    `json:"place_of_birth" db:"place_of_birth"`
	DateOfBirth  time.Time `json:"date_of_birth" db:"date_of_birth"`
	Salary       float64   `json:"salary" db:"salary"`
	PhotoKTP     string    `json:"photo_ktp" db:"photo_ktp"`
	PhotoSelfie  string    `json:"photo_selfie" db:"photo_selfie"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
