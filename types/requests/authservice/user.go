package authservice

// UserRequest ...
type UserRequest struct {
	UserName string `json:"user_name"`
	//Telephone string `json:"telephone"`
	Password string `json:"password"`
}

// UserListRequest 用户列表请求体
type UserListRequest struct {
	UserName   string `form:"user_name" binding:"required"`                  // 用户名
	OrderField string `form:"order_field" binding:"omitempty,oneof=id"`      // 排序字段
	OrderType  string `form:"order_type" binding:"omitempty,oneof=asc desc"` // 排序,降序/升序
	Page       int    `form:"page"`                                          // 页数
	PageSize   int    `form:"page_size"`                                     // 页大小
}

// UserRoleRequest 分配用户角色请求体
type UserRoleRequest struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}
