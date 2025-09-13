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
	requestDuration    metric.Float64Histogram
	itemsInCart        metric.Int64Gauge
)

func InitMeters() {
	meter := otel.Meter(serviceName)
	var err error

	successfulRequests, err = meter.Int64Counter(
		"successful_requests_total",
		metric.WithDescription("Total number of successful requests"),
	)
	if err != nil {
		log.Fatalf("failed to create successful_requests_total counter: %v", err)
	}
	failedRequests, err = meter.Int64Counter(
		"failed_requests_total",
		metric.WithDescription("Total number of failed requests"),
	)
	if err != nil {
		log.Fatalf("failed to create failed_requests_total counter: %v", err)
	}
	requestDuration, err = meter.Float64Histogram(
		"request_latency",
		metric.WithDescription("Measures the latency of a HTTP method"),
		metric.WithUnit("ms"),
	)
	if err != nil {
		log.Fatalf("failed to create request_latency histogram: %v", err)
	}
	itemsInCart, err = meter.Int64Gauge(
		"items_in_cart",
		metric.WithDescription("Current no. of items in a given cart"),
	)
	if err != nil {
		log.Fatalf("failed to create items_in_cart guage: %v", err)
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

func RecordLatency(ctx context.Context, duration float64, method string, endpoint string) {
	requestDuration.Record(ctx, duration, metric.WithAttributes(
		attribute.String("http.method", method),
		attribute.String("http.endpoint", endpoint),
	))
}

func RecordItemsInCart(ctx context.Context, cartName string, quantity int) {
	itemsInCart.Record(ctx, int64(quantity), metric.WithAttributes(
		attribute.String("cart_name", cartName),
	))
}
