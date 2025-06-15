package dto

type CreateTransaction struct {
	TenorMonth    int    `json:"tenor_month" validate:"required,oneof=1 2 3 4"`
	OTRPrice      int    `json:"otr_price" validate:"required,gt=0"`
	AdminFee      int    `json:"admin_fee" validate:"required"`
	Installment   int    `json:"installment" validate:"required"`
	Interest      int    `json:"interest" validate:"required"`
	AssetName     string `json:"asset_name" validate:"required"`
	SourceChannel string `json:"source_channe" validate:"-"`
}
