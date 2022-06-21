package api

import (
	"net/http"

	"github.com/cpw0321/mammoth/types/responses"

	"github.com/gin-gonic/gin"
)

// GetVersion 获取版本号
func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, responses.Response{
		Code:    0,
		Message: "success",
		Data:    "v1.0.0",
	})
}
