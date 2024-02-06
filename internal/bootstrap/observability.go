package bootstrap

import (
	"context"

	"github.com/Dosugamea/go-todo-api-otel/internal/infrastructure/observability"
	"github.com/Dosugamea/go-todo-api-otel/internal/infrastructure/observability/tracer_provider"
)

func InitTracer(jaegerEndpoint string, envName string) {
	ctx := context.Background()
	resource := observability.GetOtelResource(ctx, envName)
	// Tracerを初期化
	tracer_provider.InitJaegerTracer(ctx, jaegerEndpoint, resource, envName == "development")
}
