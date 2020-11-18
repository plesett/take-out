package routers

import (
	"github.com/gin-gonic/gin"
	"rely/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.Tnit)

	// 用户路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user", controller.CreateUser)
	}

	return r
}