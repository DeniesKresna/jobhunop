package Controllers

import (
	"github.com/DeniesKresna/jobhunop/Configs"
	"github.com/DeniesKresna/jobhunop/Models"
	"github.com/DeniesKresna/jobhunop/Response"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func RoleIndex(c *gin.Context) {
	var roles []Models.Role
	if err := Configs.DB.Preload("Creator").Scopes(Models.Paginate(c)).Find(&roles).Error; err != nil {
		Response.Json(c, 404, "No Role data")
	} else {
		Response.Json(c, 200, roles)
	}
}

func RoleStore(c *gin.Context) {
	var role Models.Role
	c.ShouldBindJSON(&role)

	v := validate.Struct(role)
	if !v.Validate() {
		Response.Json(c, 422, v.Errors.One())
		return
	}

	sessionId, _ := c.Get("sessionId")
	role.CreatorID = sessionId.(uint)

	if err := Configs.DB.Model(Models.Role{}).Create(&role).Error; err != nil {
		Response.Json(c, 500, "Cant Create Role")
	} else {
		Response.Json(c, 200, "Success")
	}
}

func RoleUpdate(c *gin.Context) {
	var roleUpdateInput Models.RoleUpdate
	c.ShouldBindJSON(&roleUpdateInput)
	v := validate.Struct(roleUpdateInput)
	if !v.Validate() {
		Response.Json(c, 422, v.Errors.One())
		return
	}

	if err := Configs.DB.First(&Models.Role{}, c.Param("id")).Updates(&roleUpdateInput).Error; err != nil {
		Response.Json(c, 500, "Cant Update Role")
	} else {
		Response.Json(c, 200, "Success")
	}
}
