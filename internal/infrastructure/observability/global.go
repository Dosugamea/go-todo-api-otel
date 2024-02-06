package observability

import (
	"context"

	"github.com/labstack/echo/v4"
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

func (t *TracerWrapper) StartInterfaceSpan(c echo.Context, methodName string) (context.Context, trace.Span) {
	ctx, span := t.tracer.Start(c.Request().Context(), methodName+" (I)", LAYER_ATTR_INTERFACE)
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
