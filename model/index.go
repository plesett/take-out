package model

import (
	"rely/dao"
)

/**
 * 建立 更新数据表
 */
func InitModel()  {
	dao.DB.AutoMigrate(
		&JxfUser{},
		&JxfAdder{},
	)
}
