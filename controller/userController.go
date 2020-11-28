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
	if username == "" && mobile == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": tool.ErrorReturnForMat(tool.ErrorMessageDefect, nil),
		})
		return
	}

	// 判断手机号是否合法
	if !tool.IsMobile(mobile) {
		c.JSON(http.StatusOK, gin.H{
			"data": tool.ErrorReturnForMat(tool.ErrorMobile, nil),
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
		"msg": tool.SuccessSubmit,
	})
}
