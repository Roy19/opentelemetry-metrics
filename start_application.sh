#!/bin/bash
make build

SERVICE_NAME=opentelemetry-demo \
    INSECURE_MODE=false \
    OTEL_EXPORTER_OTLP_HEADERS=signoz-ingestion-key=<signoz ingestion key> \
    OTEL_EXPORTER_OTLP_ENDPOINT=ingest.<region>.signoz.cloud:443 \
    ./app