package Routers

import (
	"github.com/DeniesKresna/jobhunop/Controllers"
	"github.com/DeniesKresna/jobhunop/Middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/", Middlewares.Auth("administrator"))

		auth.GET("/users", Controllers.UserIndex)
		auth.POST("/users", Controllers.UserStore)
		auth.PUT("/users/:id", Controllers.UserUpdate)

		auth.GET("/roles", Controllers.RoleIndex)
		auth.POST("/roles", Controllers.RoleStore)
		auth.PUT("/roles/:id", Controllers.RoleUpdate)

		v1.POST("users/login", Controllers.AuthLogin)

		//v1.GET("users", Controllers.UserIndex)
		//v1.GET("users/:id", Controllers.ShowUser)
		//v1.PUT("users/:id", Controllers.UserUpdate)
		//v1.DELETE("users/:id", Controllers.DestroyUser)
	}
	return r
}
