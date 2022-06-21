package service

import (
	"github.com/cpw0321/mammoth/api"
	v1 "github.com/cpw0321/mammoth/api/v1"
)

// RegisterRouters ...
func RegisterRouters(s *Server) {
	g := s.Engine

	// 静态资源
	g.Static("/static", "./static")

	r := g.Group("/api/v1")

	r.GET("/version", api.GetVersion) // 获取版本号

	// 用户相关
	user := v1.NewUser()
	r.POST("/user/wx/login", user.Login) // 用户登录
}
