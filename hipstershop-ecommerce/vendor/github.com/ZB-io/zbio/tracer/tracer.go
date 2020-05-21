/*
Package tracer is a wrapper over opentelemetry trace package.

*/
package tracer

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ZB-io/zbio/tracer/otel"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
)

// Current traceLib in use
const traceLib string = "otel"

type TraceCfg struct {
	Name string
}

type Tracer struct {
	tracer trace.Tracer
}

type Span struct {
	span   trace.Span
	tracer *Tracer
}

// Tracerr interface, currently Tracer struct is there so not exposing as interface
// TODO: Change name to tracer, As currently struct with same name is defined, so can't change
type Tracerr interface {
	Start(ctx context.Context, spanName string, opts ...StartOption) (context.Context, Span)
	WithSpan(
		ctx context.Context,
		spanName string,
		fn func(ctx context.Context) error,
	) error
}

// StartOption applies changes to StartConfig that sets options at span start time.
type StartOption func(*StartConfig)

// StartConfig provides options to set properties of span at the time of starting
// a new span.
type StartConfig struct {
	Attributes []core.KeyValue
	StartTime  time.Time
	Links      []Link
	Record     bool
	NewRoot    bool
	SpanKind   SpanKind
}

// Link to another spanContext
type Link struct {
	core.SpanContext
	Attributes []core.KeyValue
}

// EndConfig provides options to set properties of span at the time of ending
// the span.
type EndConfig struct {
	EndTime time.Time
}

// EndOption applies changes to EndConfig that sets options when the span is ended.
type EndOption func(*EndConfig)

// Spann interface, currently Span struct is there so not exposing as interface
type Spann interface {
	// Tracer returns tracer used to create this span. Tracer cannot be nil.
	Tracer() Tracer

	// End completes the span. No updates are allowed to span after it
	// ends. The only exception is setting status of the span.
	End(options ...EndOption)

	// AddEvent adds an event to the span.
	AddEvent(ctx context.Context, msg string, attrs ...core.KeyValue)
	// AddEventWithTimestamp adds an event with a custom timestamp
	// to the span.
	AddEventWithTimestamp(ctx context.Context, timestamp time.Time, msg string, attrs ...core.KeyValue)
}

// TraceKind assist in creating unique trace name of specified kind
type TraceKind string

var (
	// TraceKindClient is used on client side trace creation
	TraceKindClient TraceKind = "client"
	// TraceKindServer is used on server side trace creation
	TraceKindServer TraceKind = "server"
	// TraceKindBroker is used on broker side trace creation
	TraceKindBroker TraceKind = "broker"
)

// SpanKind assist in creating unique trace name of specified kind
type SpanKind string

var (
	// SpanKindClient is used on client side trace creation
	SpanKindClient SpanKind = "client"
	// SpanKindServer is used on server side trace creation
	SpanKindServer SpanKind = "server"
	// SpanKindBroker is used on broker side trace creation
	SpanKindBroker SpanKind = "broker"
)

// TraceOptions accepts
// 		- Name (required): name of tracer which would be prefixed with zbio notations
// 		- Kind (required): type of TraceKind
// 		- UUID future addition to accept unique id as an input instead of system generated unique ID
type TraceOptions struct {
	Name string
	Kind TraceKind
}

// NewTracer is factory for available tracing libs
func NewTracer(opt TraceOptions) otel.Tracer {
	// Return one of various implementations of tracer
	return otel.NewTracer(tracerName(opt.Name, opt.Kind))
}

// tracerName returns unique name when creating new trace. If name is empty, nano timestamp is postfixed
func tracerName(name string, kind TraceKind) string {
	if name == "" {
		name = strconv.FormatInt(time.Now().Unix(), 10)
	}
	return fmt.Sprintf("%s.%s.%s", "zbio.trace", kind, name)
}

// SpanFromContext eturns span interface
func SpanFromContext(ctx context.Context) otel.Span {
	if traceLib == "otel" {
		return otel.SpanFromContext(ctx)
	}
	return nil
}

// WithNewRoot ignores reference to parent trace and creates new trace ID.
func WithNewRoot() StartOption {
	return func(c *StartConfig) {
		c.NewRoot = true
	}
}

// WithSpanKind specifies the role a Span on a Trace.
func WithSpanKind(sk SpanKind) StartOption {
	return func(c *StartConfig) {
		c.SpanKind = sk
	}
}

// WithAttributes sets multiple attributes to span preserving order
func WithAttributes(attrs ...core.KeyValue) StartOption {
	return func(c *StartConfig) {
		c.Attributes = append(c.Attributes, attrs...)
	}
}

// New tracer
func New(cfg TraceCfg) *Tracer {
	tracer := Tracer{
		tracer: global.TraceProvider().Tracer(cfg.Name),
	}
	return &tracer
}

func (tr *Tracer) StartNewSpan(ctx context.Context, spanName string) (context.Context, Span) {

	var span Span
	ctx, span.span = tr.tracer.Start(ctx, spanName, trace.WithNewRoot())
	return ctx, span
}

func (tr *Tracer) StartSpanWithParent(ctx context.Context, spanName string, parentSpan *Span, kv []core.KeyValue) (context.Context, Span) {

	span := Span{tracer: tr}

	ctx, span.span = tr.tracer.Start(ctx, spanName, trace.LinkedTo(parentSpan.span.SpanContext()))
	if kv != nil {
		span.AddEventWithTimestamp(ctx, spanName, kv)
	}
	return ctx, span
}

func (span *Span) Tracer() *Tracer {
	return span.tracer
}

func (span *Span) End() {
	span.span.End()
}

func (tr *Tracer) SpanFromContext(ctx context.Context) *Span {

	var span Span
	if span.span = trace.SpanFromContext(ctx); span.span == nil {
		return nil
	}
	return &span
}

func (span *Span) AddEventWithTimestamp(ctx context.Context, eName string, kv []core.KeyValue) {
	span.span.AddEventWithTimestamp(ctx, time.Now(), eName, kv...)
}
