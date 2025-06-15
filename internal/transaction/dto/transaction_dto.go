package dto

import "time"

type CreateTransaction struct {
	TenorMonth    int    `json:"tenor_month" validate:"required,oneof=1 2 3 4"`
	OTRPrice      int    `json:"otr_price" validate:"required,gt=0"`
	AdminFee      int    `json:"admin_fee" validate:"required"`
	Installment   int    `json:"installment" validate:"required"`
	Interest      int    `json:"interest" validate:"required"`
	AssetName     string `json:"asset_name" validate:"required"`
	SourceChannel string `json:"source_channe" validate:"-"`
}
type GetTransactionHistoryResponse struct {
	ContractNumber string    `db:"contract_number" json:"contract_number"`
	TenorMonth     int       `db:"tenor_month" json:"tenor_month"`
	OTRPrice       int       `db:"otr_price" json:"otr_price"`
	Installment    int       `db:"installment" json:"installment"`
	Interest       int       `db:"interest" json:"interest"`
	AssetName      string    `db:"asset_name" json:"asset_name"`
	SourceChannel  string    `db:"source_channel" json:"source_channel"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}

type GetTransactionHistoryQuery struct {
	Page       int    `query:"page" json:"page"`
	Limit      int    `query:"limit" json:"limit"`
	ConsumerId string `query:"consumerId" json:"consumerId"`
}
