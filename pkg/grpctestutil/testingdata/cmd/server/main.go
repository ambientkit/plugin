package main

import (
	"context"
	"errors"

	"net/http"
	"time"

	stdlog "log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/grpctestutil"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// name is the Tracer name used to identify this instrumentation library.
const name = "fib"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, log, err := grpctestutil.StandardSetup(true)
	//app, err := grpctestutil.GRPCSetup(true)
	if err != nil {
		stdlog.Fatalln(err.Error())
	}

	// c := newOT(log)
	// defer c()

	// c := newOT2(log)
	// defer c.Close(context.Background())

	tp, f, err := newOT3(log, "http://localhost:14268/api/traces", "ambient", "dev", 123)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer f(ctx)

	app.PluginSystem().SetGrant("hello", ambient.GrantPluginNeighborGrantRead)

	h, err := app.Handler()
	if err != nil {
		log.Fatal(err.Error())
	}

	h = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Each execution of the run loop, we should get a new "root" span and context.
			ctx, span := tp.Tracer(name).Start(context.Background(), "Middleware")
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

// func newOT(log ambient.AppLogger) func() {
// 	exp, err := newExporter(&LogWriter{log})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	tp := trace.NewTracerProvider(
// 		trace.WithBatcher(exp),
// 		trace.WithResource(newResource()),
// 	)
// 	otel.SetTracerProvider(tp)
// 	return func() {
// 		log.Error("Hit shutdown")
// 		if err := tp.Shutdown(context.Background()); err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	}
// }

// func newOT2(log ambient.AppLogger) trace.Provider {
// 	//ctx := context.Background()

// 	// Bootstrap tracer.
// 	prv, err := trace.NewProvider(trace.ProviderConfig{
// 		JaegerEndpoint: "http://localhost:14268/api/traces",
// 		ServiceName:    "client",
// 		ServiceVersion: "1.0.0",
// 		Environment:    "dev",
// 		Disabled:       false,
// 	})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	//defer prv.Close(ctx)

// 	return prv
// }

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func newOT3(log ambient.AppLogger, url string, service string, environment string, id int64) (*trace.TracerProvider, func(ctx context.Context), error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, nil, err
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	// Cleanly shutdown and flush telemetry when the application exits.
	f := func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err.Error())
		}
	}

	return tp, f, nil
}

// // newResource returns a resource describing this application.
// func newResource() *resource.Resource {
// 	r, _ := resource.Merge(
// 		resource.Default(),
// 		resource.NewWithAttributes(
// 			semconv.SchemaURL,
// 			semconv.ServiceNameKey.String("fib"),
// 			semconv.ServiceVersionKey.String("v0.1.0"),
// 			attribute.String("environment", "demo"),
// 		),
// 	)
// 	return r
// }

// // LogWriter .
// type LogWriter struct {
// 	ambient.AppLogger
// }

// // Write .
// func (l *LogWriter) Write(p []byte) (n int, err error) {
// 	l.Error(string(p))
// 	return len(p), nil
// }

// // newExporter returns a console exporter.
// func newExporter(w io.Writer) (trace.SpanExporter, error) {
// 	return stdouttrace.New(
// 		stdouttrace.WithWriter(w),
// 		// Use human-readable output.
// 		//stdouttrace.WithPrettyPrint(),
// 		// Do not print timestamps for the demo.
// 		stdouttrace.WithoutTimestamps(),
// 	)
// }
