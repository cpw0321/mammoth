package service

import (
	"fmt"

	"github.com/cpw0321/mammoth/services/service/middleware"

	"github.com/cpw0321/mammoth/config"
	"github.com/cpw0321/mammoth/logger"
	"github.com/gin-gonic/gin"
)

// Server gin server
type Server struct {
	Engine *gin.Engine
}

// NewServer 构造函数
func NewServer() Server {
	return Server{
		gin.New(),
	}
}
func (s Server) Middleware() {
	g := s.Engine
	g.Use(middleware.Cors())
	g.Use(middleware.Logger())
	g.Use(middleware.Recovery())
}

func (s Server) RegisterRouters() {
	RegisterRouters(&s)
}

// Start 封装gin默认Start, 后期可增加功能
func (s Server) Start() error {
	gin.SetMode(config.Conf.Server.RunMode)
	s.Middleware()
	s.RegisterRouters()
	conf := config.Conf
	addr := fmt.Sprintf("%s:%d", conf.Server.Address, conf.Server.Port)
	logger.Log.Infof("server running addr: %v", addr)

	g := s.Engine
	return g.Run(addr)
}

// Stop 虚增加优雅停止
func (s Server) Stop() {}
