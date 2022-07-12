package authservice

// ListOfRoleResponseBody 用户列表返回体
type ListOfRoleResponseBody struct {
	Total    int64      `json:"total"`    // 总记录数
	Page     int        `json:"page"`     // 页码
	PageSize int        `json:"pageSize"` // 页大小
	List     []RoleBody `json:"list"`     // 列表
}

type RoleBody struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
