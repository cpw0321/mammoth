package authservice

import (
	"gorm.io/gorm"
)

// User 用户表结构
type User struct {
	gorm.Model

	UserName string `json:"user_name,omitempty" gorm:"column:user_name; type:varchar(255); comment:'用户名称'"` // 用户真实名称
	Password string `json:"password,omitempty" gorm:"column:password; type:varchar(255); comment:'用户密码'"`   // 用户密码
	//Telephone string `json:"telephone" gorm:"column:telephone; type:varchar(20); comment:'手机号码'"`            // 手机号码
}
