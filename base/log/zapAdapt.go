package log

/**

@Author argus.deng


日志框架，底层使用的时 uber的zap，
当前使用了lumberjack库对文件进行切割，
软件cloud 基础功能实现后，可以使用 可以用ELK、kafka等文件聚合中间件代替


-- 配置了两个输出源/file
  log level debug 模式时 all enable
  log level > debug  only enable file write

how to use

use fun var logger = log.GetLogger() in any where

https://blog.csdn.net/weixin_34144848/article/details/91371535
*/

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var (
	logger   *zap.Logger
	loglevel = "info"
)

func GetLogger() *zap.Logger {
	return logger
}

func init() {

	log.Print("init log module")

	checkEnvAndInitZap()
}

func checkEnvAndInitZap() {
	//
	logger = initZap("./log.log", loglevel)
}

func initZap(logpath string, loglevel string) *zap.Logger {

	var allCore []zapcore.Core
	var level zapcore.Level

	switch loglevel {

	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	//console begin
	cw := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	//console over

	// file log begin
	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    128,     // megabytes
		MaxBackups: 30,      // 最多保留300个备份
		MaxAge:     7,       // days
		Compress:   true,    // 是否压缩 disabled by default
	}
	defer hook.Close()

	w := zapcore.AddSync(&hook)

	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	fcore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	allCore = append(allCore, fcore)
	// file log over

	if loglevel == "debug" {
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, cw, level))
	}

	coreT := zapcore.NewTee(allCore...)

	logger = zap.New(coreT).WithOptions(zap.AddCaller())

	logger.Info("Default Logger init success", zap.String("log level", level.String()))

	return logger
}
