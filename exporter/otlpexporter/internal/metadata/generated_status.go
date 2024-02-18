// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

var (
	Type      = component.MustNewType("otlp")
	scopeName = "go.opentelemetry.io/collector/exporter/otlpexporter"
)

const (
	LogsStability    = component.StabilityLevelBeta
	TracesStability  = component.StabilityLevelStable
	MetricsStability = component.StabilityLevelStable
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter(scopeName)
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer(scopeName)
}
