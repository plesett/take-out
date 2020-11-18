package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"rely/model"
	"rely/tool"
)

/**
 * 创建用户路由
 */
func CreateUser(c *gin.Context)  {

	username := c.PostForm("username")
	mobile := c.PostForm("mobile")

	// 判断是否存在用户名
	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg": "信息请填写完整",
			"status": false,
		})
		return
	}

	// 判断手机号是否合法
	if !tool.IsMobile(mobile) {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg": "手机号码异常",
			"status": false,
		})
		return
	}

	var userDto = model.JxfUser{
		Uuid:     tool.UniqueId(),
		Username: username,
		Mobile:   mobile,
		Address:  "",
		Model:    gorm.Model{},
	}

	if err := model.CreateAUser(&userDto); err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
