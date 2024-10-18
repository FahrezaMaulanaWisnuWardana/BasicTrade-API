package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UUID      string    `gorm:"not null"`
	Name      string    `gorm:"not null" form:"name" json:"name" valid:"required~Name of product is required"`
	ImageUrl  string    `gorm:"not null" form:"image" json:"image" valid:"required~Image of product is required"`
	AdminID   uint      `gorm:"index"`
	Admin     *Admin    `gorm:"foreignKey:AdminID"`
	Variants  []Variant `gorm:"foreignKey:ProductID"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
