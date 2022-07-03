package go_zaplogger_iso8601

import (
	zap "go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

type myZapLogger struct {
	myLogger *zap.SugaredLogger
}

func InitZapLogger(logger *zap.Logger) (Logger, error) {
	sugaredLogger := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()

	return &zapLogger{
		sugaredLogger: sugaredLogger,
	}, nil
}

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

	initialLogger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	logger := initialLogger.Sugar()
	defer logger.Sync()

	logger.Info("Logger init successful.")

	if warnInvalidLevel {
		logger.Warn("Invalid value provided for logLevel. Valid values are: 'debug', 'info', 'warn', 'error'.")
	}

	//return logger
	return &myZapLogger{
		myLogger: logger,
	}
}



var logger Logger

type Logger interface {
	Debug(args ...interface{})

	Info(args ...interface{})

	Warn(args ...interface{})

	Error(args ...interface{})

	Panic(args ...interface{})

	Fatal(args ...interface{})
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}
func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}


func (l *myZapLogger) Debug(args ...interface{}) {
	l.myLogger.Debug(args...)
}

func (l *myZapLogger) Info(args ...interface{}) {
	l.myLogger.Info(args...)
}

func (l *myZapLogger) Warn(args ...interface{}) {
	l.myLogger.Warn(args...)
}
func (l *myZapLogger) Error(args ...interface{}) {
	l.myLogger.Error(args...)
}

func (l *myZapLogger) Panic(args ...interface{}) {
	l.myLogger.Panic(args...)
}

func (l myZapLogger) Fatal(args ...interface{}) {
	l.myLogger.Fatal(args...)
}