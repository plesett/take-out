package main

import (
	"fmt"
	"rely/dao"
	"rely/model"
)

func main() {
	if err := dao.InitMySQL(); err != nil {
		panic("数据库连接失败")
	}
	defer dao.Close()
	// 创建数据库表
	model.AutoMigrateDB()
	fmt.Println("结束")
}
