package go_zaplogger_iso8601

import (
	zap "go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
	"errors"
	"strings"
	"fmt"
)

var logger Logger

type Logger interface {
	Debug(args ...interface{})

	Info(args ...interface{})

	Warn(args ...interface{})

	Error(args ...interface{})

	Panic(args ...interface{})

	Fatal(args ...interface{})
}

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
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

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	if strings.HasPrefix(caller.Function, "github.com") {
		parts := strings.Split(caller.Function, "/")
		functionName := parts[len(parts)-1]
		enc.AppendString(fmt.Sprintf("%s:%d - %s", caller.TrimmedPath(), caller.Line, functionName))
	}
}


func InitLogger(filePath string, logLevel string) (Logger, error) {

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
		EncodeCaller:     customCallerEncoder,
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

	var invalidLevelErr error
	if warnInvalidLevel {
		invalidLevelErr = errors.New("invalid value provided for logLevel. Defaulting to 'info'")
	}

	logger := initialLogger.Sugar()
	defer logger.Sync()

	return &zapLogger{
		sugaredLogger: logger,
	},
	invalidLevelErr
}

