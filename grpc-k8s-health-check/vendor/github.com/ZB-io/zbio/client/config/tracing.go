package config

import (
	"github.com/ZB-io/zbio/log"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/exporters/trace/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// InitTrace sets OpenTelemetry STDOUT && Jaeger exporter with SDKTrace provider
func InitTrace() {
	// Stdout Exporter
	_, err := stdout.NewExporter(stdout.Options{PrettyPrint: true})
	if err != nil {
		log.Fatalf("Error encountered aquiring exporter: %v", err)
	}

	// Jaeger Exporter
	jaegerExporter, err := jaeger.NewRawExporter(
		jaeger.WithCollectorEndpoint("http://jaeger-collector:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "client",
			Tags: []core.KeyValue{
				key.String("exporter", "jaeger"),
				key.String("version", "zbio/zbclient:v1"),
			},
		}),
	)
	if err != nil {
		log.Fatalf("Error encountered acquiring Jaeger Exporter: %v", err)
	}
	tp, err := sdktrace.NewProvider(
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		// sdktrace.WithSyncer(exporter),
		sdktrace.WithSyncer(jaegerExporter),
	)
	if err != nil {
		log.Fatalf("Error encountered setting sdktrace provider: %v", err)
	}
	global.SetTraceProvider(tp)
}
