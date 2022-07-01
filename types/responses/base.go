package responses

import (
	"net/http"

	"github.com/cpw0321/mammoth/internal"
)

// Response 返回体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功
func Success(c *internal.ServiceContext, data interface{}) {
	c.Context.JSON(http.StatusOK, Response{
		0,
		"success",
		data,
	})
}

// Fail 失败
func Fail(c *internal.ServiceContext, code int, msg string) {
	c.Context.JSON(http.StatusOK, Response{
		code,
		msg,
		"",
	})
}
