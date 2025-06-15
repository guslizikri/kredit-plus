package dto

type SetLimit struct {
	TenorMonth  int `json:"tenor_month" db:"tenor_month" validate:"required,oneof=1 2 3 4"`
	LimitAmount int `json:"limit_amount" db:"limit_amount" validate:"required"`
}
