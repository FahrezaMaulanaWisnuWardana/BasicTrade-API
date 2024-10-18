package request

type AdminRequest struct {
	UUID     string `gorm:"not null"`
	Name     string `gorm:"not null" form:"name" json:"name" valid:"required~Name is required"`
	Email    string `gorm:"not null" form:"email" json:"email" valid:"required~Email is required"`
	Password string `gorm:"not null" form:"password" json:"password" valid:"required~Password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
}
