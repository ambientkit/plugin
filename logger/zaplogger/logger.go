package zaplogger

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/ambientkit/ambient"
	"github.com/mattn/go-colorable"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
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

	tracerProvider *sdktrace.TracerProvider
	ctx            context.Context
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

func (l *Logger) output(f1 func(args ...interface{}), f2 func(template string, args ...interface{}), loglevel string, format string, v ...interface{}) {
	name := l.serviceName
	if len(name) == 0 {
		name = l.appName
	}

	format, v = l.service(format, v...)
	if len(format) == 0 {
		f1(v...)
		if l.ctx != nil {
			_, file, line, _ := runtime.Caller(2)
			if strings.Contains(file, "zaplogger") {
				_, file, line, _ = runtime.Caller(3)
			}
			_, span := l.tracerProvider.Tracer(name).Start(l.ctx, fmt.Sprint(v...))
			span.SetAttributes(attribute.String("log.level", loglevel))
			span.SetAttributes(attribute.String("log.message", fmt.Sprint(v...)))
			span.SetAttributes(attribute.String("caller.file", file))
			span.SetAttributes(attribute.Int("caller.line", line))
			if strings.EqualFold(os.Getenv("AMB_TRACE_STACK"), "true") {
				span.SetAttributes(attribute.String("stack", string(debug.Stack())))
			}
			span.End()
		}
	} else {
		f2(format, v...)
		if l.ctx != nil {
			_, file, line, _ := runtime.Caller(2)
			if strings.Contains(file, "zaplogger") {
				_, file, line, _ = runtime.Caller(3)
			}
			_, span := l.tracerProvider.Tracer(name).Start(l.ctx, fmt.Sprintf(format, v...))
			span.SetAttributes(attribute.String("log.level", loglevel))
			span.SetAttributes(attribute.String("log.message", fmt.Sprintf(format, v...)))
			span.SetAttributes(attribute.String("caller.file", file))
			span.SetAttributes(attribute.Int("caller.line", line))
			if strings.EqualFold(os.Getenv("AMB_TRACE_STACK"), "true") {
				span.SetAttributes(attribute.String("stack", string(debug.Stack())))
			}
			span.End()
		}
	}
}

// Debug is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Debug(format string, v ...interface{}) {
	l.output(l.logentry().Debug, l.logentry().Debugf, "debug", format, v...)
}

// Info is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Info(format string, v ...interface{}) {
	l.output(l.logentry().Info, l.logentry().Infof, "info", format, v...)
}

// Warn is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Warn(format string, v ...interface{}) {
	l.output(l.logentry().Warn, l.logentry().Warnf, "warn", format, v...)
}

// Error is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Error(format string, v ...interface{}) {
	l.output(l.logentry().Error, l.logentry().Errorf, "error", format, v...)
}

// Fatal is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty. It's followed by a call
// to os.Exit(1).
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.output(l.logentry().Fatal, l.logentry().Fatalf, "fatal", format, v...)
}

// Name returns the name of the logger.
func (l *Logger) Name() string {
	return l.appName
}

// Named returns a new logger with the appended name, linked to the existing
// logger.
func (l *Logger) Named(serviceName string) ambient.AppLogger {
	out := l.clone()
	out.serviceName = serviceName
	return out
}

// clone returns a copy of the logger.
func (l *Logger) clone() *Logger {
	out := &Logger{
		appName:        l.appName,
		appVersion:     l.appVersion,
		serviceName:    l.serviceName,
		log:            l.log,
		loglevel:       l.loglevel,
		tracerProvider: l.tracerProvider,
	}

	return out
}

// For returns a context-aware logger to support OpenTracing.
func (l *Logger) For(ctx context.Context) ambient.Logger {
	if span := trace.SpanFromContext(ctx); span != nil {
		logger := l.clone()
		logger.ctx = ctx
		// TODO: Determine if these need to be saved.
		//span.SpanContext().TraceID()
		//span.SpanContext().SpanID()
		return logger
	}
	return l
}

// SetTracerProvider sets the OpenTelemetry tracer provider.
func (l *Logger) SetTracerProvider(tp *sdktrace.TracerProvider) {
	l.tracerProvider = tp
}

// Trace returns a context and span to support OpenTracing.
func (l *Logger) Trace(ctx context.Context, spanName string) (context.Context, trace.Span) {
	name := l.serviceName
	if len(name) == 0 {
		name = l.appName
	}
	return l.tracerProvider.Tracer(name).Start(ctx, spanName)
}
