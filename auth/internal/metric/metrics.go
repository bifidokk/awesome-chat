package metric

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	namespace = "auth_space"
	appName   = "auth"
)

// Metrics is a structure that holds Prometheus metrics for monitoring purposes.
type Metrics struct {
	requestCounter        prometheus.Counter
	responseCounter       *prometheus.CounterVec
	histogramResponseTime *prometheus.HistogramVec
}

var metrics *Metrics

// Init initializes the Prometheus metrics for monitoring gRPC services.
func Init(_ context.Context) {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_requests_total",
				Help:      "Total requests",
			},
		),
		responseCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_responses_total",
				Help:      "Total responses",
			},
			[]string{"status", "method"},
		),
		histogramResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_histogram_response_time_seconds",
				Help:      "Server response time in seconds",
				Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
			},
			[]string{"status"},
		),
	}
}

// IncRequestCounter increments the total request counter metric.
func IncRequestCounter() {
	metrics.requestCounter.Inc()
}

// IncResponseCounter increments the response counter metric based on the status and method labels.
func IncResponseCounter(status string, method string) {
	metrics.responseCounter.WithLabelValues(status, method).Inc()
}

// HistogramResponseTimeObserve records the observed response time in the histogram metric
func HistogramResponseTimeObserve(status string, time float64) {
	metrics.histogramResponseTime.WithLabelValues(status).Observe(time)
}
