package model

import "time"

type Limit struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	TenorMonths int       `db:"tenor_months"`
	LimitAmount int       `db:"limit_amount"`
	UsedAmount  int       `db:"used_amount"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
