package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// defaultLogPath 默认日志输出路径
const defaultLogPath = "logs" // 默认输出日志文件路径

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// initLogger 初始化 log
func initLogger(levelStr, logPath string) error {
	writeSyncer, err := getLogWriter(logPath) // 日志文件配置 文件位置和切割
	if err != nil {
		return err
	}
	encoder := getEncoder() // 获取日志输出编码
	level, ok := levelMap[levelStr]
	if !ok {
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller()) // zap.Addcaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	// 1. zap.ReplaceGlobals 函数将当前初始化的 logger 替换到全局的 logger,
	// 2. 使用 logger 的时候 直接通过 zap.S().Debugf("xxx") or zap.L().Debug("xxx")
	// 3. 使用 zap.S() 和 zap.L() 提供全局锁，保证一个全局的安全访问logger的方式
	zap.ReplaceGlobals(logger)
	//zap.L().Debug("")
	//zap.S().Debugf("")
	return nil
}

// getEncoder 编码器(如何写入日志)
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // log 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 输出level序列化为全大写字符串，如 INFO DEBUG ERROR
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
}

// getLogWriter 获取日志输出方式  日志文件 控制台
func getLogWriter(logPath string) (zapcore.WriteSyncer, error) {

	// 判断日志路径是否存在，如果不存在就创建
	if exist := isExist(logPath); !exist {
		if logPath == "" {
			logPath = defaultLogPath
		}
		if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
			fmt.Printf("日志目录{ %s }创建失败,  将要使用默认的日志目录: { %s }, 本次错误信息: %v,\n", logPath, defaultLogPath, err)
			logPath = defaultLogPath
			if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}

	// 日志文件 与 日志切割 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logPath, "out"), // 日志文件路径
		MaxSize:    100,                           // 单个日志文件最大多少 mb
		MaxBackups: 5,                             // 日志备份数量
		MaxAge:     30,                            // 日志最长保留时间
		Compress:   false,                         // 是否压缩日志
		LocalTime:  true,
	}

	// 日志只输出到日志文件
	return zapcore.AddSync(lumberJackLogger), nil
}

// isExist 判断文件或者目录是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Load 加载日志配置
func Load(level, logPath string) error {
	if err := initLogger(level, logPath); err != nil {
		return err
	}

	initInstance()

	return nil
}
