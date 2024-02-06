package observability

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

func GetOtelResource(ctx context.Context, deployEnvName string) *resource.Resource {
	// 常にSpanに付与させるリソースを設定
	resource, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			attribute.String("service.name", "go-todo-api-otel"),
			attribute.String("library.language", "go"),
			attribute.String("deployment.environment", deployEnvName),
		))
	if err != nil {
		log.Fatal(err)
	}
	return resource
}
