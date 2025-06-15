package model

import "time"

type Limit struct {
	ID          string    `db:"id"`
	ConsumerID  string    `db:"consumer_id"`
	TenorMonth  int       `db:"tenor_month"`
	LimitAmount int       `db:"limit_amount"`
	UsedAmount  int       `db:"used_amount"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
