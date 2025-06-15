package model

import "time"

type Transaction struct {
	ID             string    `db:"id"`
	ConsumerID     string    `db:"consumer_id"`
	LimitID        string    `db:"limit_id"`
	ContractNumber string    `db:"contract_number"`
	OTRPrice       int       `db:"otr_price"`
	AdminFee       int       `db:"admin_fee"`
	Installment    int       `db:"installment"`
	Interest       int       `db:"interest"`
	AssetName      string    `db:"asset_name"`
	SourceChannel  string    `db:"source_channel"`
	CreatedAt      time.Time `db:"created_at"`
}
