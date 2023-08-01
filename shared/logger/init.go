package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	cfg "tuzi-tiktok/config"
)

var (
	logger Logger
	level  zap.AtomicLevel
)

const development = false
const (
	jsonEncoding    = "json"
	consoleEncoding = "console"
)

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

func getLogLevel(level string) zapcore.Level {
	if zapLevel, ok := levelMap[strings.ToLower(level)]; ok {
		return zapLevel
	}
	return zapcore.InfoLevel
}

func SetLogLevel(l string) {
	level.SetLevel(getLogLevel(l))
}

func loadConfig(c *zap.Config) {
	SetLogLevel(cfg.LoggerConfig.Level)
	c.Development = cfg.LoggerConfig.Development
	c.Encoding = consoleEncoding
	encoding := cfg.LoggerConfig.Encoding
	if encoding == jsonEncoding || encoding == consoleEncoding {
		c.Encoding = encoding
	}
}

// init Default Logger
func init() {
	c := zap.NewDevelopmentConfig()
	level = zap.NewAtomicLevelAt(zap.DebugLevel)
	c.Level = level
	loadConfig(&c)

	c.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapLogger, err := c.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger = &CLogger{zapLogger.Sugar()}
}

func GetLogger() Logger {
	return logger
}
