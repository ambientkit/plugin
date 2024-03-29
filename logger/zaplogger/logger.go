package zaplogger

import (
	"fmt"
	"io"

	"github.com/ambientkit/ambient"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger represents a logger.
type Logger struct {
	log      *zap.SugaredLogger
	loglevel *zap.AtomicLevel

	appName     string
	appVersion  string
	serviceName string
}

// NewLogger returns a new logger with a default log level of error.
func NewLogger(appName string, appVersion string, optionalWriter io.Writer) *Logger {
	loglevel := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "" // Disable timestamps.
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var writer io.Writer
	if optionalWriter == nil {
		writer = colorable.NewColorableStdout()
	} else {
		writer = optionalWriter
	}

	base := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		//zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(writer),
		loglevel,
	))

	defer base.Sync()
	sugar := base.Sugar()

	return &Logger{
		log:      sugar,
		loglevel: &loglevel,

		appName:    appName,
		appVersion: appVersion,
	}
}

func (l *Logger) logentry() *zap.SugaredLogger {
	return l.log.Named(l.appName + " v" + l.appVersion)
}

// SetLogLevel will set the logger output level.
func (l *Logger) SetLogLevel(level ambient.LogLevel) {
	// Set log level temporarily to info.
	l.loglevel.SetLevel(zap.InfoLevel)

	var loglevel zapcore.Level

	switch level {
	case ambient.LogLevelDebug:
		loglevel = zapcore.DebugLevel
		l.logentry().Infof("zaplogger: log level set to: %v", "debug")
	case ambient.LogLevelInfo:
		loglevel = zapcore.InfoLevel
		l.logentry().Infof("zaplogger: log level set to: %v", "info")
	case ambient.LogLevelWarn:
		loglevel = zapcore.WarnLevel
		l.logentry().Infof("zaplogger: log level set to: %v", "warn")
	case ambient.LogLevelError:
		loglevel = zapcore.ErrorLevel
		l.logentry().Infof("zaplogger: log level set to: %v", "error")
	case ambient.LogLevelFatal:
		loglevel = zapcore.FatalLevel
		l.logentry().Infof("zaplogger: log level set to: %v", "fatal")
	default:
		loglevel = zapcore.InfoLevel
		l.logentry().Infof("zaplogger: log level set to: %v", "info")
	}

	l.loglevel.SetLevel(loglevel)
}

// Log is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Log(level ambient.LogLevel, format string, v ...interface{}) {
	switch level {
	case ambient.LogLevelDebug:
		l.Debug(format, v...)
	case ambient.LogLevelInfo:
		l.Info(format, v...)
	case ambient.LogLevelWarn:
		l.Warn(format, v...)
	case ambient.LogLevelError:
		l.Error(format, v...)
	case ambient.LogLevelFatal:
		l.Fatal(format, v...)
	default:
		l.Info(format, v...)
	}
}

func (l *Logger) service(format string, v ...interface{}) (string, []interface{}) {
	if len(l.serviceName) == 0 {
		return format, v
	}

	if len(format) == 0 {
		return format, append([]interface{}{l.serviceName}, v...)
	}
	return fmt.Sprintf("%v: %v", l.serviceName, format), v
}

// Debug is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Debug(format string, v ...interface{}) {
	format, v = l.service(format, v...)
	if len(format) == 0 {
		l.logentry().Debug(v...)
	} else {
		l.logentry().Debugf(format, v...)
	}
}

// Info is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Info(format string, v ...interface{}) {
	format, v = l.service(format, v...)
	if len(format) == 0 {
		l.logentry().Info(v...)
	} else {
		l.logentry().Infof(format, v...)
	}
}

// Warn is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Warn(format string, v ...interface{}) {
	format, v = l.service(format, v...)
	if len(format) == 0 {
		l.logentry().Warn(v...)
	} else {
		l.logentry().Warnf(format, v...)
	}
}

// Error is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Error(format string, v ...interface{}) {
	format, v = l.service(format, v...)
	if len(format) == 0 {
		l.logentry().Error(v...)
	} else {
		l.logentry().Errorf(format, v...)
	}
}

// Fatal is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty. It's followed by a call
// to os.Exit(1).
func (l *Logger) Fatal(format string, v ...interface{}) {
	format, v = l.service(format, v...)
	if len(format) == 0 {
		l.logentry().Fatal(v...)
	} else {
		l.logentry().Fatalf(format, v...)
	}
}

// Name returns the name of the logger.
func (l *Logger) Name() string {
	return l.appName
}

// Named returns a new logger with the appended name, linked to the existing
// logger.
func (l *Logger) Named(serviceName string) ambient.AppLogger {
	return &Logger{
		appName:     l.appName,
		appVersion:  l.appVersion,
		serviceName: serviceName,
		log:         l.log,
		loglevel:    l.loglevel,
	}
}
