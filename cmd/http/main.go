// Copyright 2022 The mammoth Authors

// Package main implement function main
package main

import (
	"flag"
	"log"

	"github.com/cpw0321/mammoth/config"
	"github.com/cpw0321/mammoth/datasource/mysql"
	"github.com/cpw0321/mammoth/logger"
	"github.com/cpw0321/mammoth/services/service"
)

var configFile = flag.String("f", "etc/config.toml", "the config file")

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("service recover err: %v", err)
		}
	}()

	flag.Parse()
	// 初始化配置文件
	config.InitConfig(*configFile)
	// 初始化日志
	logger.InitLogger()
	// 连接数据库
	mysql.InitDB()
	// 初始化服务
	svc := service.NewServer()
	svc.Start()
}
