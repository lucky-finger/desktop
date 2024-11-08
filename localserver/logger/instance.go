package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var instance *zap.Logger

func initInstance() {
	instance = zap.L()
	Debug = instance.Debug
	Info = instance.Info
	Warn = instance.Warn
	Error = instance.Error
	DPanic = instance.DPanic
	Panic = instance.Panic
	Fatal = instance.Fatal
	With = instance.With
	Level = instance.Level
}

func Instance() Logger {
	return instance
}

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(string, ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	With(fields ...zap.Field) *zap.Logger
	Level() zapcore.Level
}

var (
	Debug  func(msg string, fields ...zap.Field)
	Info   func(msg string, fields ...zap.Field)
	Warn   func(msg string, fields ...zap.Field)
	Error  func(msg string, fields ...zap.Field)
	DPanic func(msg string, fields ...zap.Field)
	Panic  func(msg string, fields ...zap.Field)
	Fatal  func(msg string, fields ...zap.Field)
	With   func(fields ...zap.Field) *zap.Logger
	Level  func() zapcore.Level
)
