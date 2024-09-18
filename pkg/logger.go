package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"Users/config"
)

func InitLogger(logInfo *config.Config) *zap.Logger {
	var logger *zap.Logger

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	var level zapcore.LevelEnabler
	switch logInfo.Logs.Level {
	case "error":
		level = zapcore.ErrorLevel
	case "info":
		level = zapcore.InfoLevel
	case "debug":
		level = zapcore.DebugLevel
	case "warning":
		level = zapcore.WarnLevel
	default:
		level = zapcore.ErrorLevel
	}

	if logInfo.Logs.Path != "" {
		fileEncoder := zapcore.NewJSONEncoder(config)
		logRotation := &lumberjack.Logger{
			Filename:   logInfo.Logs.Path,
			MaxBackups: logInfo.Logs.MaxBackups,
			MaxAge:     logInfo.Logs.MaxAge,
		}

		writer := zapcore.AddSync(logRotation)

		core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, level))
		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		return logger
	}

	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
