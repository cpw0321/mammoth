// Copyright 2022 The mammoth Authors

// Package middleware implements middleware
package middleware

import (
	"github.com/cpw0321/mammoth/logger"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.Request.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = xid.New().String()
		}
		// 请求IP
		clientIP := c.ClientIP()
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqURI := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()

		logger.Log.Infof("reqID: %v clientIP: %v reqMethod: %v reqURI: %v statusCode: %v", reqID, clientIP, reqMethod, reqURI, statusCode)
		//contentType := c.Request.Header.Get("Content-Type")
		//if strings.Contains(strings.ToLower(contentType), "application/json") {
		//	// 获取body中请求参数
		//	data, err := c.GetRawData()
		//	if err != nil {
		//		logger.Log.Errorf("get request body is failed! err: %v", err)
		//	}
		//	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		//	// 日志格式
		//	logger.Log.Infof("reqID: %v clientIP: %v reqMethod: %v reqURI: %v statusCode: %v reqBody: %v", reqID, clientIP, reqMethod, reqURI, statusCode, string(data))
		//} else {
		//	logger.Log.Infof("reqID: %v clientIP: %v reqMethod: %v reqURI: %v statusCode: %v", reqID, clientIP, reqMethod, reqURI, statusCode)
		//}

		// 处理请求
		c.Next()
	}
}
