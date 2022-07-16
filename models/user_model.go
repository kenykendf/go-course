package models

import (
	"errors"
	"final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)"`
	Age      uint   `gorm:"not null" json:"age" form:"age" valid:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil

	if u.Age < 8 {
		return errors.New("please mindly back to registration after you are 9 years old")
	}
	return
}
