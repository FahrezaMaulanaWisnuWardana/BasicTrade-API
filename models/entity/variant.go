package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Variant struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UUID        string `gorm:"not null"`
	VariantName string `gorm:"type:varchar(255);not null" json:"variant_name"`
	Quantity    int    `gorm:"not null"`
	ProductID   uint   `json:"product_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(v)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
