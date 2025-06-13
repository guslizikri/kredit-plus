package dto

type CreateLimit struct {
	UserID      string `json:"user_id" db:"user_id" validate:"required"`
	TenorMonths int    `json:"tenor_months" db:"tenor_months" validate:"required"`
	LimitAmount int    `json:"limit_amount" db:"limit_amount" validate:"required"`
	UsedAmount  int    `json:"used_amount" db:"used_amount" validate:"required"`
}
