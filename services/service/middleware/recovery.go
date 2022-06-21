// Copyright 2022 The mammoth Authors

// Package middleware implements middleware
package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/cpw0321/mammoth/logger"
	"github.com/cpw0321/mammoth/types/responses"

	"github.com/gin-gonic/gin"
)

// Recovery panic中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover(); err != nil {
			reqID := c.Request.Header.Get("X-Request-ID")
			// 请求IP
			clientIP := c.ClientIP()
			// 请求方式
			reqMethod := c.Request.Method
			// 请求路由
			reqURI := c.Request.RequestURI
			// 状态码
			statusCode := c.Writer.Status()
			// 获取body中请求参数
			data, err := c.GetRawData()
			if err != nil {
				logger.Log.Errorf("get request body is failed! err: %v", err)
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			// 日志格式
			logger.Log.Errorf("service panic, reqID: %v clientIP: %v reqMethod: %v reqURI: %v statusCode: %v reqBody: %v err: %v", reqID, clientIP, reqMethod, reqURI, statusCode, string(data), err)

			c.JSON(http.StatusInternalServerError, responses.Response{})
		}

		// 处理请求
		c.Next()
	}
}
