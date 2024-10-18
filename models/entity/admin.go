package entity

import (
	"BasicTrade-API/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UUID      string `gorm:"not null"`
	Name      string `gorm:"not null" form:"name" json:"name" valid:"required~Name is required"`
	Email     string `gorm:"not null" form:"email" json:"email" valid:"required~Email is required"`
	Password  string `gorm:"not null" form:"password" json:"password" valid:"required~Password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(a)

	if errCreate != nil {
		err = errCreate
		return
	}
	a.Password = helpers.HasPass(a.Password)

	err = nil
	return
}
