package requests

// UserRequest ...
type UserRequest struct {
	UserName string `json:"user_name"`
	//Telephone string `json:"telephone"`
	Password string `json:"password"`
}

// UserListRequest 用户列表请求体
type UserListRequest struct {
	UserName   string `form:"user_name" binding:"required"`                  // 用户名
	Type       string `form:"type" binding:"required"`                       // 密码
	OrderField string `form:"order_field" binding:"omitempty,oneof=id"`      // 排序字段
	OrderType  string `form:"order_type" binding:"omitempty,oneof=ASC DESC"` // 排序,降序/升序
	Page       int    `form:"page"`                                          // 页数
	PageSize   int    `form:"page_size"`                                     // 页大小
}
