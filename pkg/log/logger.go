package log

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LogOption ...
type LogOption struct {
	// StdOut is true will print log to stdout
	StdOut bool
	// Filename is the file to write logs to, if empty, log will output to stdout
	LogPath string
	// log level
	Level zapcore.Level

	// MaxAge is the maximum number of days to retain old log files
	MaxAge int
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int

	// MaxBackups is the maximum number of old log files to retain.
	MaxBackups int
}

// NewLogger ...
func NewLogger(option LogOption) logr.Logger {
	if option.Level == 0 {
		option.Level = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), getLogWriter(option), option.Level)
	logger := zap.New(core, zap.AddCaller())
	return zapr.NewLogger(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(option LogOption) zapcore.WriteSyncer {
	if len(option.LogPath) > 0 {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   option.LogPath,
			MaxSize:    option.MaxSize,
			MaxBackups: option.MaxBackups,
			MaxAge:     option.MaxAge,
		}
		if option.StdOut {
			return zapcore.NewMultiWriteSyncer(
				zapcore.AddSync(lumberJackLogger),
				zapcore.AddSync(os.Stdout),
			)
		}
		return zapcore.AddSync(lumberJackLogger)
	}

	return zapcore.AddSync(os.Stdout)
}
