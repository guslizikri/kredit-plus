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
