package observability

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var LAYER_ATTR_INTERFACE = trace.WithAttributes(
	attribute.String("internal.layer", "interface"),
)

var LAYER_ATTR_USECASE = trace.WithAttributes(
	attribute.String("internal.layer", "usecase"),
)

var LAYER_ATTR_PERSISTENCE = trace.WithAttributes(
	attribute.String("internal.layer", "persistence"),
)
