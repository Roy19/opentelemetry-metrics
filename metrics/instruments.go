package metrics

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var (
	SuccessfulRequests metric.Int64Counter
	FailedRequests     metric.Int64Counter
)

func InitMeters() {
	meter := otel.Meter(serviceName)
	var err error

	SuccessfulRequests, err = meter.Int64Counter(
		"successful_requests_total",
		metric.WithDescription("Total number of successful requests"),
	)
	if err != nil {
		log.Fatalf("Failed to create successful_requests_total counter: %v", err)
	}
	FailedRequests, err = meter.Int64Counter(
		"failed_requests_total",
		metric.WithDescription("Total number of failed requests"),
	)
	if err != nil {
		log.Fatalf("Failed to create failed_requests_total counter: %v", err)
	}
}
