package logruslogger

import (
	"context"
	"io"

	"github.com/ambientkit/ambient"
	"github.com/sirupsen/logrus"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// Logger represents a logger.
type Logger struct {
	log *logrus.Logger

	appName     string
	appVersion  string
	serviceName string

	tracerProvider *sdktrace.TracerProvider
	span           trace.Span
}

// NewLogger returns a new logger with a default log level of error.
func NewLogger(appName string, appVersion string, optionalWriter io.Writer) *Logger {
	var base = logrus.New()
	//base.SetFormatter(&logrus.JSONFormatter{})
	base.Level = logrus.InfoLevel
	if optionalWriter != nil {
		base.Out = optionalWriter
	}

	return &Logger{
		log: base,

		appName:    appName,
		appVersion: appVersion,
	}
}

func (l *Logger) logentry() *logrus.Entry {
	standardFields := logrus.Fields{
		"app":     l.appName,
		"version": l.appVersion,
	}

	if len(l.serviceName) > 0 {
		standardFields["service"] = l.serviceName
	}

	return l.log.WithFields(standardFields)
}

// SetLogLevel will set the logger output level.
func (l *Logger) SetLogLevel(level ambient.LogLevel) {
	// Set log level temporarily to info.
	l.log.Level = logrus.InfoLevel
	loglevel := logrus.InfoLevel

	switch level {
	case ambient.LogLevelDebug:
		loglevel = logrus.DebugLevel
		l.logentry().Infoln("logruslogger: log level set to:", "debug")
	case ambient.LogLevelInfo:
		loglevel = logrus.InfoLevel
		l.logentry().Infoln("logruslogger: log level set to:", "info")
	case ambient.LogLevelWarn:
		loglevel = logrus.WarnLevel
		l.logentry().Infoln("logruslogger: log level set to:", "warn")
	case ambient.LogLevelError:
		loglevel = logrus.ErrorLevel
		l.logentry().Infoln("logruslogger: log level set to:", "error")
	case ambient.LogLevelFatal:
		loglevel = logrus.FatalLevel
		l.logentry().Infoln("loglogrusloggerrus: log level set to:", "fatal")
	default:
		loglevel = logrus.InfoLevel
		l.logentry().Infoln("logruslogger: log level set to:", "info")
	}

	l.log.Level = loglevel
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

// Debug is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Debug(format string, v ...interface{}) {
	if len(format) == 0 {
		l.logentry().Debugln(v...)
	} else {
		l.logentry().Debugf(format, v...)
	}
}

// Info is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Info(format string, v ...interface{}) {
	if len(format) == 0 {
		l.logentry().Infoln(v...)
	} else {
		l.logentry().Infof(format, v...)
	}
}

// Warn is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Warn(format string, v ...interface{}) {
	if len(format) == 0 {
		l.logentry().Warnln(v...)
	} else {
		l.logentry().Warnf(format, v...)
	}
}

// Error is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty.
func (l *Logger) Error(format string, v ...interface{}) {
	if len(format) == 0 {
		l.logentry().Errorln(v...)
	} else {
		l.logentry().Errorf(format, v...)
	}
}

// Fatal is equivalent to log.Printf() + "\n" if format is not empty.
// It's equivalent to Println() if format is empty. It's followed by a call
// to os.Exit(1).
func (l *Logger) Fatal(format string, v ...interface{}) {
	if len(format) == 0 {
		l.logentry().Fatalln(v...)
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
		tracerProvider: l.tracerProvider,
	}

	return out
}

// For returns a context-aware logger to support OpenTracing.
func (l *Logger) For(ctx context.Context) ambient.Logger {

	if span := trace.SpanFromContext(ctx); span != nil {
		logger := l.clone()
		logger.span = span

		//sc := span.SpanContext()
		//span.SetAttributes()

		// if jaegerCtx, ok := sc.(jaeger.SpanContext); ok {
		// 	logger.spanFields = []zapcore.Field{
		// 		zap.String("trace_id", jaegerCtx.TraceID().String()),
		// 		zap.String("span_id", jaegerCtx.SpanID().String()),
		// 	}
		// }

		return logger
	}
	return l
}

// SetTracerProvider sets the tracer provider.
func (l *Logger) SetTracerProvider(tp *sdktrace.TracerProvider) {
	l.tracerProvider = tp
}

// Trace returns a context and an OpenTelemetry span.
func (l *Logger) Trace(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return l.tracerProvider.Tracer(l.appName).Start(ctx, spanName)
}
