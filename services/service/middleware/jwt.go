// Copyright 2022 The mammoth Authors

// Package middleware implements middleware
package middleware

import (
	"log"
	"net/http"

	"github.com/cpw0321/mammoth/utils"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    1000,
				"message": "无权限访问",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		// parseToken 解析token包含的信息
		claims, err := utils.ParseToken(token)
		if err != nil {
			if err == utils.ErrTokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code":    1000,
					"message": "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
