package Models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"size:256" validate:"required"`
	Password string `json:"password" form:"password" gorm:"size:256" validate:"required"`
	Email    string `json:"email" form:"email" gorm:"size:256" validate:"required,email"`
	RoleId   uint8  `json:"role_id"`
}

func (b *User) TableName() string {
	return "users"
}

var validate = validator.New()

func (b *User) ValidateData() (errorMessage string, ok bool) {
	err := validate.Struct(b)
	errorMessage = ""
	ok = false
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage = errorMessage + err.StructField() + " Input Wrong, "
			ok = true
			return
		}
	}
	return
}
