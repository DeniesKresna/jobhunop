package Routers

import (
	"github.com/DeniesKresna/jobhunop/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("users", Controllers.UserIndex)
		//v1.GET("users/:id", Controllers.ShowUser)
		v1.POST("users", Controllers.UserStore)
		//v1.PUT("users/:id", Controllers.UpdateUser)
		//v1.DELETE("users/:id", Controllers.DestroyUser)
	}
	return r
}
