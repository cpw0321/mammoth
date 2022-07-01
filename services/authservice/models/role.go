package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	Name      string // 角色名称
	Status    int    // 状态 1-启用 2-禁用
	CreatorId int    // 创建者id
}
