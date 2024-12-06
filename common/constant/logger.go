package constant

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LoggerTypeDev  = "development"
	LoggerTypeProd = "production"
)

const (
	LoggerDebug = "debug"
	LoggerInfo  = "info"
	LoggerWarn  = "warnning"
	LoggerError = "error"
	LoggerPanic = "panic"
	LoggerFatal = "fatal"
)

const (
	LoggerEncodingJSON    = "json"
	LoggerEncodingConsole = "console"
)

const (
	LoggerPath = "logs/simulator.log"
)

var ZapLoggerMap = map[string]zapcore.Level{
	LoggerDebug: zap.DebugLevel,
	LoggerInfo:  zap.InfoLevel,
	LoggerWarn:  zap.WarnLevel,
	LoggerError: zap.ErrorLevel,
	LoggerPanic: zap.PanicLevel,
	LoggerFatal: zap.FatalLevel,
}
