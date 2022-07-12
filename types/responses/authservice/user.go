package authservice

// UserLoginResponseBody 用户登录返回体
type UserLoginResponseBody struct {
	Token string `json:"token"` // 用户token
}

// ListOfUserResponseBody 用户列表返回体
type ListOfUserResponseBody struct {
	Total    int64      `json:"total"`    // 总记录数
	Page     int        `json:"page"`     // 页码
	PageSize int        `json:"pageSize"` // 页大小
	List     []UserBody `json:"list"`     // 列表
}

type UserBody struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
}
