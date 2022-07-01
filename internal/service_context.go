package internal

import "github.com/gin-gonic/gin"

// ServiceContext 定制服务上下文
type ServiceContext struct {
	Context *gin.Context
}

type HandlerFunc func(*ServiceContext)

// Handler 转换context
func Handler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(ServiceContext)
		context.Context = c
		handler(context)
	}
}
