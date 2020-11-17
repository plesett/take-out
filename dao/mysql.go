package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

/**
 * 连接mysql 数据库
 */
func InitMySQL()(err error)  {
	dsn := "db_test:zbBLMwZsSMnpeZCD@(47.101.206.144:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}

/**
 * 关闭数据库连接
 */
func Close()  {
	DB.Close()
}