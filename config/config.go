// Copyright 2022 The mammoth Authors

// Package config implement config
package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf *Config

// Config 配置结构
type Config struct {
	Server struct {
		Address    string
		Port       int    // 端口
		HasPprof   bool   // 是否启用pprof查看性能
		HasSwagger bool   //  是否展示swagger api
		RunMode    string // 运行模式debug/release
	}
	Mysql struct {
		Host          string
		Port          int
		DBName        string
		User          string
		Password      string
		MaxOpenConns  int
		MaxidleConns  int
		DBTablePrefix string
	}
	Log struct {
		Filename   string // 日志输出文件路径
		MaxSize    int    // 文件大小 单位兆字节, 默认是 100M
		MaxBackups int    // 最多保留备份个数, 默认是 3
		ToFile     bool   // 是否输出到Stdout, 默认不输出到Stdout
		IsJSON     bool   // 是否按JSON格式输出，默认不按JSON输出
		Level      string // 日志级别
		MaxAge     int    // 文件最多保存多少天
		Compress   bool   // 是否压缩
	}
	Redis struct {
		Addr     string
		Username string
		Password string
		DB       int
	}
	Rabbitmq struct {
		Addr string
	}
	Elasticsearch struct {
		Urls []string
	}
}

// InitConfig 初始化配置文件
func InitConfig(configFile string) {
	viper.SetConfigFile(configFile)
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: ", err)
	}
	// https://github.com/spf13/viper#unmarshaling
	// 解析配置文件
	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
