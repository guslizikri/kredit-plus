package dto

type ConsumerLogin struct {
	NIK      string `json:"nik" db:"nik" validate:"required"`
	FullName string `json:"full_name" db:"full_name" validate:"required"`
}

type AdminLogin struct {
	Username string `json:"username" db:"username" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
}
