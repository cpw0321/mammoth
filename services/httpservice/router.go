package service

import (
	"github.com/cpw0321/mammoth/api"
	v1 "github.com/cpw0321/mammoth/api/v1"
	"github.com/cpw0321/mammoth/internal"
)

// RegisterRouters ...
func RegisterRouters(s *Server) {
	g := s.Engine

	// 静态资源
	g.Static("/static", "./static")

	r := g.Group("/api/v1")

	r.GET("/version", internal.Handler(api.GetVersion)) // 获取版本号

	// 认证服务相关
	auth := v1.NewAuth()
	r.POST("/user/register", internal.Handler(auth.Register)) // 用户注册
	r.POST("/user/login", internal.Handler(auth.Login))       // 用户登录
	r.GET("/user/list", internal.Handler(auth.UserList))      // 获取用户列表
	r.POST("/user/role", internal.Handler(auth.UserRole))     // 分配用户角色

	r.POST("/role/create", internal.Handler(auth.RoleCreate)) // 新建角色
	r.GET("/role/list", internal.Handler(auth.RoleList))      // 获取角色列表
	//ag := r.Use(middleware.Casbin())
}
