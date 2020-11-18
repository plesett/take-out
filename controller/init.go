package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 项目初始化 提示
 */
func Tnit(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"msg": "项目已成功启动！",
	})
}