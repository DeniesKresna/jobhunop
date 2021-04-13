package Helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidateData(c *gin.Context, entity interface{}) {
	err := validate.Struct(entity)
	errorMessage := ""
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage = errorMessage + err.StructField()
			RespondJSON(c, 422, "invalid input "+errorMessage)
		}
	}
}
