package dto

import "time"

type CreateConsumer struct {
	NIK         string    `form:"nik" db:"nik" validate:"required"`
	FullName    string    `form:"full_name" db:"full_name" validate:"required"`
	LegalName   string    `form:"legal_name" db:"legal_name" validate:"required"`
	BirthPlace  string    `form:"birth_place" db:"birth_place" validate:"-"`
	BirthDate   time.Time `form:"birth_date" db:"birth_date" time_format:"2006-01-02" validate:"-"`
	Salary      int       `form:"salary" db:"salary" validate:"required"`
	PhotoKTP    string    `form:"-" db:"photo_ktp" validate:"-"`
	PhotoSelfie string    `form:"-" db:"photo_selfie" validate:"-"`
}
type LimitEmbedded struct {
	TenorMonth      int `json:"tenor_month" db:"tenor_month"`
	LimitAmount     int `json:"limit_amount" db:"limit_amount"`
	UsedAmount      int `json:"used_amount" db:"used_amount"`
	AvailableAmount int `json:"available_amount" db:"available_amount"`
}
type GetConsumerDetailResponse struct {
	ID          string          `json:"id" db:"id"`
	NIK         string          `json:"nik" db:"nik"`
	FullName    string          `json:"full_name" db:"full_name"`
	LegalName   string          `json:"legal_name" db:"legal_name"`
	BirthPlace  string          `json:"birth_place" db:"birth_place"`
	BirthDate   time.Time       `json:"birth_date" db:"birth_date"`
	Salary      int             `json:"salary" db:"salary"`
	PhotoKTP    string          `json:"photo_ktp" db:"photo_ktp"`
	PhotoSelfie string          `json:"photo_selfie" db:"photo_selfie"`
	Limits      []LimitEmbedded `json:"limits" db:"limits"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" db:"updated_at"`
}
