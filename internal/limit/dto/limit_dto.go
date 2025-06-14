package dto

type SetLimit struct {
	TenorMonth  int `json:"tenor_month" db:"tenor_month" validate:"required"`
	LimitAmount int `json:"limit_amount" db:"limit_amount" validate:"required"`
}
