// Copyright 2022 The mammoth Authors

// Package logger implement log
package logger

import (
	"os"

	"github.com/cpw0321/mammoth/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var levelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
	"panic": zap.PanicLevel,
	"fatal": zap.FatalLevel,
}

// Log 全局日志
var Log *Logger

// Logger ...
type Logger struct {
	*zap.SugaredLogger
}

// InitLogger 日常使用的log, 建议使用json效率更高
func InitLogger() {
	conf := config.Conf
	hook := lumberjack.Logger{
		Filename:   conf.Log.Filename,   // 日志文件路径
		MaxSize:    conf.Log.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: conf.Log.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     conf.Log.MaxAge,     // 文件最多保存多少天
		Compress:   conf.Log.Compress,   // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 是否输出到控制台
	var syncer zapcore.WriteSyncer
	if conf.Log.ToFile {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	} else {
		syncer = zapcore.AddSync(&hook)
	}

	// 是否json格式输出
	var encoder zapcore.Encoder
	if conf.Log.IsJSON {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(levelMap[conf.Log.Level])

	core := zapcore.NewCore(
		encoder,     // 编码器配置
		syncer,      // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 构造日志
	logger := zap.New(core, caller, development)

	Log = &Logger{
		logger.Sugar(),
	}
}
