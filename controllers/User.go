Package Controllers

import (
	"github.com/DeniesKresna/jobhunop/Configs"
	"github.com/DeniesKresna/jobhunop/Helpers"
	"github.com/DeniesKresna/jobhunop/Models"
	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	var users []Models.User
	if err := Configs.DB.Scopes(Helpers.Paginate(c)).Find(&users).Error; err != nil {
		Helpers.RespondJSON(c, 404, users)
	} else {
		Helpers.RespondJSON(c, 200, users)
	}
}

func UserStore(c *gin.Context) {
	var user Models.User
	c.Bind(&user)

	errString, ok := user.ValidateData()

	if ok {
		Helpers.RespondJSON(c, 422, errString)
		return
	}

	hashedPassword, err := Helpers.Hash(user.Password)
	if err != nil {
		Helpers.RespondJSON(c, 400, "error hashing password")
		return
	}
	user.Password = string(hashedPassword)
	user.RoleId = 1

	if err := Configs.DB.Model(Models.User{}).Create(&user).Error; err != nil {
		Helpers.RespondJSON(c, 404, user)
	} else {
		Helpers.RespondJSON(c, 200, user)
	}
}

func UserUpdate(c *gin.Context){
	var user Model.User
	c.Bind(&user)

	
}
