package request

type VariantRequest struct {
	VariantName string `gorm:"type:varchar(255);not null" form:"variant_name"`
	Quantity    int    `gorm:"not null" form:"quantity"`
	ProductUUID string `gorm:"type:uuid;not null" form:"product_id"`
}
