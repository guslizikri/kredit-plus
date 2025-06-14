package dto

type CreateLimit struct {
	ConsumerID  string `json:"consumer_id" db:"consumer_id" validate:"required"`
	TenorMonths int    `json:"tenor_months" db:"tenor_months" validate:"required"`
	LimitAmount int    `json:"limit_amount" db:"limit_amount" validate:"required"`
	UsedAmount  int    `json:"used_amount" db:"used_amount" validate:"required"`
}
