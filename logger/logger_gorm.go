// Copyright 2022 The mammoth Authors

// Package logger 适配 gorm logger
// 参考 https://github.com/mpalmer/gorm-zerolog
package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"gorm.io/gorm/logger"
)

type gormLogger struct {
	log *Logger
}

// NewGormLogger 根据 zerolog Logger 构建 gorm logger.Interface
func NewGormLogger(log *Logger) logger.Interface {
	return gormLogger{
		log: log,
	}
}

func (gl gormLogger) LogMode(logger.LogLevel) logger.Interface {
	return gl
}

func (gl gormLogger) Error(ctx context.Context, msg string, opts ...interface{}) {
	gl.log.Error(fmt.Sprintf(msg, opts...))
}

func (gl gormLogger) Warn(ctx context.Context, msg string, opts ...interface{}) {
	gl.log.Warn(fmt.Sprintf(msg, opts...))
}

func (gl gormLogger) Info(ctx context.Context, msg string, opts ...interface{}) {
	gl.log.Info(fmt.Sprintf(msg, opts...))
}

func (gl gormLogger) Trace(ctx context.Context, begin time.Time, f func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := f()
	if err != nil {
		gl.log.Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	} else {
		gl.log.Debug("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql)) // zerolog.Level 为debug时就能输出所有sql日志了
	}
}
