package tracer

import (
	"context"
	"os"
	"time"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/plugin/grpctrace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor intercepts and extracts incoming trace data
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	requestMetadata, _ := metadata.FromIncomingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	entries, spanCtx := grpctrace.Extract(ctx, &metadataCopy)
	ctx = correlation.ContextWithMap(ctx, correlation.NewMap(correlation.MapUpdate{
		MultiKV: entries,
	}))

	hostname, _ := os.Hostname()
	serverSpanAttrs := []core.KeyValue{
		key.New("zbio.server.hostname").String(hostname),
	}

	tr := global.TraceProvider().Tracer("zbio.service.server.interceptor.tracer")
	ctx, span := tr.Start(
		trace.ContextWithRemoteSpanContext(ctx, spanCtx),
		"zbio.service.server.interceptor.span",
		trace.WithStartTime(time.Now()),
		trace.WithAttributes(serverSpanAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	return handler(ctx, req)
}

// UnaryClientInterceptor intercepts and injects outgoing trace
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	requestMetadata, _ := metadata.FromOutgoingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	tr := global.TraceProvider().Tracer("zbio.client.interceptor.tracer")
	err := tr.WithSpan(ctx, "zbio.client.interceptor.span",
		func(ctx context.Context) error {
			grpctrace.Inject(ctx, &metadataCopy)
			ctx = metadata.NewOutgoingContext(ctx, metadataCopy)

			err := invoker(ctx, method, req, reply, cc, opts...)
			setTraceStatus(ctx, err)
			return err
		})
	return err
}

func setTraceStatus(ctx context.Context, err error) {
	if err != nil {
		s, _ := status.FromError(err)
		trace.SpanFromContext(ctx).SetStatus(s.Code(), "Error in Unary client request invocation")
	}
}

// UnaryBrokerClientInterceptor intercepts and injects outgoing trace for zbbroker
func UnaryBrokerClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	requestMetadata, _ := metadata.FromOutgoingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	tr := global.TraceProvider().Tracer("zbio.broker.client.interceptor.tracer")
	err := tr.WithSpan(ctx, "zbio.broker.client.interceptor.span",
		func(ctx context.Context) error {
			grpctrace.Inject(ctx, &metadataCopy)
			ctx = metadata.NewOutgoingContext(ctx, metadataCopy)

			err := invoker(ctx, method, req, reply, cc, opts...)
			setTraceStatus(ctx, err)
			return err
		})
	return err
}

// UnaryBrokerServerInterceptor intercepts and extracts incoming trace data on zbbroker
func UnaryBrokerServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	requestMetadata, _ := metadata.FromIncomingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	entries, spanCtx := grpctrace.Extract(ctx, &metadataCopy)
	ctx = correlation.ContextWithMap(ctx, correlation.NewMap(correlation.MapUpdate{
		MultiKV: entries,
	}))

	hostname, _ := os.Hostname()
	serverSpanAttrs := []core.KeyValue{
		key.New("zbio.server.hostname").String(hostname),
	}

	tr := global.TraceProvider().Tracer("zbio.broker.server.interceptor.tracer")
	ctx, span := tr.Start(
		trace.ContextWithRemoteSpanContext(ctx, spanCtx),
		"zbio.broker.server.interceptor.span",
		trace.WithStartTime(time.Now()),
		trace.WithAttributes(serverSpanAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	return handler(ctx, req)
}
