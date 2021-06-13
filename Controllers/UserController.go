package Controllers

import (
	"github.com/DeniesKresna/jobhunop/Configs"
	"github.com/DeniesKresna/jobhunop/Helpers"
	"github.com/DeniesKresna/jobhunop/Models"
	"github.com/DeniesKresna/jobhunop/Response"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func UserIndex(c *gin.Context) {
	var users []Models.User
	if err := Configs.DB.Preload("Role").Scopes(Models.Paginate(c)).Find(&users).Error; err != nil {
		Response.Json(c, 404, "No User data")
	} else {
		Response.Json(c, 200, users)
	}
}

func UserStore(c *gin.Context) {
	var user Models.User
	c.ShouldBindJSON(&user)

	v := validate.Struct(user)
	if !v.Validate() {
		Response.Json(c, 422, v.Errors.One())
		return
	}

	hashedPassword, err := Helpers.Hash(user.Password)
	if err != nil {
		Response.Json(c, 400, "error hashing password")
		return
	}
	user.Password = string(hashedPassword)
	user.RoleID = 1

	if err := Configs.DB.Model(Models.User{}).Create(&user).Error; err != nil {
		Response.Json(c, 500, "Cant Create User")
	} else {
		Response.Json(c, 200, "Success")
	}
}

func UserUpdate(c *gin.Context) {
	var userUpdateInput Models.UserUpdate
	c.ShouldBindJSON(&userUpdateInput)
	v := validate.Struct(userUpdateInput)
	if !v.Validate() {
		Response.Json(c, 422, v.Errors.One())
		return
	}

	if err := Configs.DB.First(&Models.User{}, c.Param("id")).Updates(&userUpdateInput).Error; err != nil {
		Response.Json(c, 500, "Cant Update User")
	} else {
		Response.Json(c, 200, "Success")
	}
}
