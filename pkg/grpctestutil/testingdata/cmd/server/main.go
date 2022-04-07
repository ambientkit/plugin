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

	tp, f, err := tracerProvider(log, "http://localhost:14268/api/traces", "ambient", "dev", 123)
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

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(log ambient.AppLogger, url string, service string, environment string, id int64) (*trace.TracerProvider, func(ctx context.Context), error) {
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
