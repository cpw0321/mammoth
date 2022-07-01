package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model

	UserId uint // 用户id
	RoleId uint // 角色id
}
