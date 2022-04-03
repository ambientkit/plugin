package main

import (
	"context"
	"errors"
	"io"
	"net/http"

	stdlog "log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/grpctestutil"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// name is the Tracer name used to identify this instrumentation library.
const name = "fib"

func main() {
	app, log, err := grpctestutil.StandardSetup(true)
	//app, err := grpctestutil.GRPCSetup(true)
	if err != nil {
		stdlog.Fatalln(err.Error())
	}

	exp, err := newExporter(&LogWriter{log})
	if err != nil {
		log.Fatal(err.Error())
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)
	defer func() {
		log.Error("Hit shutdown")
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal(err.Error())
		}
	}()
	otel.SetTracerProvider(tp)

	app.PluginSystem().SetGrant("hello", ambient.GrantPluginNeighborGrantRead)

	h, err := app.Handler()
	if err != nil {
		log.Fatal(err.Error())
	}

	h = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Each execution of the run loop, we should get a new "root" span and context.
			ctx, span := otel.Tracer(name).Start(context.Background(), "Middleware")
			defer span.End()
			log.Error("Hit middleware")
			span.SetAttributes(attribute.String("request.n", "cool"))
			span.SetStatus(codes.Ok, "it worked")
			span.RecordError(errors.New("there was an error"))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}(h)

	log.Fatal("%v", app.ListenAndServe(h))
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

// LogWriter .
type LogWriter struct {
	ambient.AppLogger
}

// Write .
func (l *LogWriter) Write(p []byte) (n int, err error) {
	l.Error(string(p))
	return len(p), nil
}

// newExporter returns a console exporter.
func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}
