// Copyright 2022 The mammoth Authors

// Package middleware implements middleware
package middleware

import (
	"log"
	"net/http"

	"github.com/cpw0321/mammoth/internal"

	"github.com/gin-gonic/gin"
)

// JWT 中间件，检查token
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "无权限访问",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		// parseToken 解析token包含的信息
		claims, err := internal.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
