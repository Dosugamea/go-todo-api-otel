package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type TracerWrapper struct {
	tracer trace.Tracer
}

func NewCustomTracer(tracer trace.Tracer) *TracerWrapper {
	return &TracerWrapper{
		tracer: tracer,
	}
}

func (t *TracerWrapper) StartInterfaceSpan(ctx context.Context, methodName string) (context.Context, trace.Span) {
	ctx, span := t.tracer.Start(ctx, methodName+" (I)", LAYER_ATTR_INTERFACE)
	return ctx, span
}

func (t *TracerWrapper) StartUsecaseSpan(ctx context.Context, methodName string) (context.Context, trace.Span) {
	ctx, span := t.tracer.Start(ctx, methodName+" (U)", LAYER_ATTR_USECASE)
	return ctx, span
}

func (t *TracerWrapper) StartPersistenceSpan(ctx context.Context, methodName string) (context.Context, trace.Span) {
	ctx, span := t.tracer.Start(ctx, methodName+" (P)", LAYER_ATTR_PERSISTENCE)
	return ctx, span
}

var Tracer = NewCustomTracer(otel.Tracer("github.com/Dosugamea/go-todo-api-otel"))
