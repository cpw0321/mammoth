package authservice

// RoleCreateRequest 创建角色请求体
type RoleCreateRequest struct {
	Name string `json:"name" binding:"required"` // 角色名
}

// RoleListRequest 角色列表请求体
type RoleListRequest struct {
	Name       string `form:"name" binding:"required"`                       // 角色名
	OrderField string `form:"order_field" binding:"omitempty,oneof=id"`      // 排序字段
	OrderType  string `form:"order_type" binding:"omitempty,oneof=asc desc"` // 排序,降序/升序
	Page       int    `form:"page"`                                          // 页数
	PageSize   int    `form:"page_size"`                                     // 页大小
}
