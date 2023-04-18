package models

import (
	"sesi_12/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your Full Name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required,email~Invalid Email Format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum of 6 charachters"`
	Level    string    `gorm:"not null" json:"level" form:"level" valid:"required~Level is required"`
	Product  []Product `gorm:"constrain:OnUpdate:CASCADE, OnDelet:SET NULL;" json:"product"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HasPass(u.Password)
	err = nil
	return
}
