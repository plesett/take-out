package main

import (
	"rely/dao"
	"rely/model"
	"rely/routers"
)

func main() {
	if err := dao.InitMySQL(); err != nil {
		panic("数据库连接失败")
	}
	defer dao.Close()
	// 创建数据库表
	model.InitModel()
	// 导入路由
	r := routers.SetupRouter()
	// 启动项目
	r.Run(":8989")
}
