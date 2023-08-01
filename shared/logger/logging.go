package logger

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})

	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
}
type CLogger struct {
	Logger
}

// Info is info level
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Warn is warning level
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Error is error level
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Debug is debug level
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Infof is format info level
func Infof(fmt string, args ...interface{}) {
	GetLogger().Infof(fmt, args...)
}

// Warnf is format warning level
func Warnf(fmt string, args ...interface{}) {
	GetLogger().Warnf(fmt, args...)
}

// Errorf is format error level
func Errorf(fmt string, args ...interface{}) {
	GetLogger().Errorf(fmt, args...)
}

// Debugf is format debug level
func Debugf(fmt string, args ...interface{}) {
	GetLogger().Debugf(fmt, args...)
}
