package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 返回体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		"success",
		data,
	})
}

// Fail 失败
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		"",
	})
}
