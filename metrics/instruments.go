package metrics

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	successfulRequests metric.Int64Counter
	failedRequests     metric.Int64Counter
)

func InitMeters() {
	meter := otel.Meter(serviceName)
	var err error

	successfulRequests, err = meter.Int64Counter(
		"successful_requests_total",
		metric.WithDescription("Total number of successful requests"),
	)
	if err != nil {
		log.Fatalf("Failed to create successful_requests_total counter: %v", err)
	}
	failedRequests, err = meter.Int64Counter(
		"failed_requests_total",
		metric.WithDescription("Total number of failed requests"),
	)
	if err != nil {
		log.Fatalf("Failed to create failed_requests_total counter: %v", err)
	}
}

func IncSuccessfulRequests(ctx context.Context, code int) {
	successfulRequests.Add(ctx, 1, metric.WithAttributes(
		attribute.Int("http.status_code", code),
	))
}

func IncFailedRequests(ctx context.Context, code int) {
	failedRequests.Add(ctx, 1, metric.WithAttributes(
		attribute.Int("http.status_code", code),
	))
}
