package otel

import (
	"context"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
)

// Tracer is zbio implentation of open telemetry
type Tracer interface {
	trace.Tracer
}

// Span use composition to inherit form upstream open telemetry
type Span interface {
	trace.Span
}

// NewTracer gives otel tracer. Pass name otherwise noopTracer would be returned
// Returning Interface as multiple types can be returned based on global tracer
func NewTracer(name string) Tracer {
	return global.TraceProvider().Tracer(name)
}

// Start accepts Tracer Interface and invokes upstream open telemetry start
func Start(ctx context.Context, tr Tracer, spanName string, opts ...trace.StartOption) {
	// TODO: Convert otel.StartOptions into trace.Startoptions
	tr.Start(ctx, spanName, opts...)
}

// SpanFromContext returns span
func SpanFromContext(ctx context.Context) Span {
	return trace.SpanFromContext(ctx)
}
