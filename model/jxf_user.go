package model

import (
	"github.com/jinzhu/gorm"
	"rely/dao"
)

/**
 * 用户表
 */
type JxfUser struct {
	Uuid uint `gorm:"not null;unique"` // 设置uuid为主键
	Username string `gorm:"not null; size: 255"` // 用户名
	Mobile string `gorm:"not null; size: 11"` // 手机号码
	Address string `gorm:"size:255"` // 收货地址
	gorm.Model
}

/**
 * 创建用户
 */
func CreateAUser(userDto *JxfUser) (err error)  {
	if err = dao.DB.Create(&userDto).Error; err != nil {
		return err
	}
	return
}

/**
 * 删除用户
 */
func DeleteAUser(userUuid string) (err error)  {
	if err = dao.DB.Unscoped().Where("uuid = ?", userUuid).Delete(&JxfUser{}).Error; err != nil {
		return err
	}
	return
}