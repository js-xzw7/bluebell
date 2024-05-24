package logger

import (
	"bluebell/settings"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init(cfg *settings.LogConfig, mode string) (err error) {

	var core zapcore.Core
	if mode == "dev" {
		//进入开发模式，日志输出终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			// zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {

		//进入生产模式，日志输出文件
		infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel || lvl == zapcore.WarnLevel
		})

		errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.ErrorLevel || lvl == zapcore.PanicLevel || lvl == zapcore.FatalLevel
		})

		accesswriteSyncer := getLogWriter(cfg.AccessLogFile, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
		errorwriteSyncer := getLogWriter(cfg.ErrorLogFile, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
		encoder := getEncoder()

		core = zapcore.NewTee(
			zapcore.NewCore(encoder, accesswriteSyncer, infoLevel),
			zapcore.NewCore(encoder, errorwriteSyncer, errorLevel),
		)
	}

	Logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Logger)
	zap.L().Info("log init success")
	return
}

func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}

	return zapcore.AddSync(lumberjackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
