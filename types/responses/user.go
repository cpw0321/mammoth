package responses

import "github.com/cpw0321/mammoth/services/authservice/models"

type ListOfUserResponseBody struct {
	Total    int64         `json:"total"`    // 总记录数
	Page     int           `json:"page"`     // 页码
	PageSize int           `json:"pageSize"` // 页大小
	List     []models.User `json:"list"`     // 列表
}
