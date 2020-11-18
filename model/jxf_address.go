package model

import "github.com/jinzhu/gorm"

/**
 * 收货地址
 */
type JxfAdder struct {
	//ID uint 			`gorm:"primary_key"`  // 主键
	UserId string		`gorm:"not null;"` // 用户 id
	ProvinceId string   `gorm:"not null;"` // 省
	CityId string 		`gorm:"not null;"` // 城市
	DistrictId string 	`gorm:"not null;"` // 区
	Name string 		`gorm:"not null"` // 收货人姓名
	Mobile string 		`gorm:"not null"` // 手机号
	Remark string       `gorm:"not null"` // 详细收货地址
	IsDefault string    `gorm:"type:enum('0', '1');default:'0'"` // 默认地址 0 - 默认 1 - 否
	gorm.Model
}
