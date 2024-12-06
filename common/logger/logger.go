package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github/leezone/simulator/common/constant"
	"github/leezone/simulator/config"
)

var Logger *ZapLogger

type ZapLogger struct {
	*zap.Logger
}

func (z ZapLogger) ErrorWithStack(msg string, err error, parameter interface{}) {
	z.Error(msg, zap.Error(err), zap.Stack("stack"), zap.Any("parameter", parameter))
}

func CreateLogger(conf config.Logger) *ZapLogger {
	core := zapcore.NewCore(
		getEncoder(conf.Type, conf.Encoding),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.Path,       // 日志文件路径
			MaxSize:    conf.MaxSize,    // 最大日志文件大小（MB），超过会切割
			MaxBackups: conf.MaxBackups, // 保留的最大备份数
			MaxAge:     conf.MaxAge,     // 保留的最大天数
			Compress:   conf.Compress,   // 是否压缩旧日志
		}),
		zap.NewAtomicLevelAt(constant.ZapLoggerMap[conf.Level]),
	)

	return &ZapLogger{
		Logger: zap.New(core),
	}
}

func getEncoder(loggerType, loggerEncoding string) zapcore.Encoder {
	var cfg zapcore.EncoderConfig
	switch loggerType {
	case constant.LoggerTypeDev:
		cfg = zap.NewDevelopmentEncoderConfig()
	default:
		cfg = zap.NewProductionEncoderConfig()
	}

	cfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)

	if loggerEncoding == constant.LoggerEncodingConsole {
		return zapcore.NewConsoleEncoder(cfg)
	}
	return zapcore.NewJSONEncoder(cfg)

}

func InitLogger(conf config.Logger) {
	Logger = CreateLogger(conf)
	Logger.Info("logger init success")
}
