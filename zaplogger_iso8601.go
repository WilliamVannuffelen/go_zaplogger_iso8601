package zaplogger_iso8601

import (
	zap "go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
)

func InitLogger(filePath string, logLevel string) (*zap.Logger, error) {

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "logger",
		CallerKey:        "caller",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " - ",
	}

	var atomicLevel zap.AtomicLevel

	switch {
	case logLevel == "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case logLevel == "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case logLevel == "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case logLevel == "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	}

	logConfig := zap.Config{
		//Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Level:            atomicLevel,
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout", filePath},
		ErrorOutputPaths: []string{"stderr", filePath},
	}

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Logger init successful.")

	return logger, err
}
