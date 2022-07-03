package go_zaplogger_iso8601

import (
	zap "go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

var logger Logger

func InitLogger(filePath string, logLevel string) Logger {

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      "function",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " - ",
	}

	atomicLevel := zap.NewAtomicLevel()
	warnInvalidLevel := false

	switch {
	case logLevel == "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case logLevel == "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case logLevel == "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case logLevel == "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	default:
		warnInvalidLevel = true
		atomicLevel.SetLevel(zap.InfoLevel)
	}
	logConfig := zap.Config{
		Level:            atomicLevel,
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout", filePath},
		ErrorOutputPaths: []string{"stderr", filePath},
	}

	initialLogger, err := logConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	logger := initialLogger.Sugar()
	defer logger.Sync()

	logger.Info("Logger init successful.")

	if warnInvalidLevel {
		logger.Warn("Invalid value provided for logLevel. Valid values are: 'debug', 'info', 'warn', 'error'.")
	}

	return &ZapLogger{
		sugaredLogger: logger,
	}
}

type Logger interface {
	Debug(args ...interface{})

	Info(args ...interface{})

	Warn(args ...interface{})

	Error(args ...interface{})

	Panic(args ...interface{})

	Fatal(args ...interface{})
}

func (l *ZapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *ZapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}
func (l *ZapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *ZapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}
