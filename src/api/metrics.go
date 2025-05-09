package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
)

// OpenCensus Tags
var (
	ProjectIDKey = tag.MustNewKey("project_id")
	HostKey      = tag.MustNewKey("host")
	ClientIPKey  = tag.MustNewKey("client_ip")
	EndpointKey  = tag.MustNewKey("endpoint")
	StatusKey    = tag.MustNewKey("status")
	MethodKey    = tag.MustNewKey("http_method")
)

// OpenCensus Measurements
var (
	TotalRequests     = stats.Int64("service/content/total_requests", "Total number of content requests", stats.UnitDimensionless)
	RequestLatency    = stats.Float64("service/content/request_latency", "Latency of HTTP requests (ms)", stats.UnitMilliseconds)
	RequestSizeBytes  = stats.Float64("service/content/request_size_bytes", "Size of HTTP requests in bytes", stats.UnitBytes)
	ResponseSizeBytes = stats.Float64("service/content/response_size_bytes", "Size of HTTP responses in bytes", stats.UnitBytes)
)

// Prometheus Metrics
var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "earthfast",
			Subsystem: "service_content",
			Name:      "total_requests",
			Help:      "Total number of content requests",
		},
		[]string{"hostname", "client_ip", "project_id", "endpoint", "host"},
	)

	latencyHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "earthfast",
			Subsystem: "service_content",
			Name:      "latency_distribution",
			Help:      "Distribution of request latencies in seconds",
			Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"hostname", "client_ip", "endpoint"},
	)

	requestSizeHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "earthfast",
			Subsystem: "service_content",
			Name:      "request_size_distribution",
			Help:      "Distribution of request sizes in bytes",
			Buckets:   prometheus.ExponentialBuckets(100, 2, 10),
		},
		[]string{"hostname", "client_ip", "endpoint"},
	)

	responseSizeHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "earthfast",
			Subsystem: "service_content",
			Name:      "response_size_distribution",
			Help:      "Distribution of response sizes in bytes",
			Buckets:   prometheus.ExponentialBuckets(100, 2, 10),
		},
		[]string{"hostname", "client_ip", "endpoint"},
	)
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(latencyHistogram)
	prometheus.MustRegister(requestSizeHistogram)
	prometheus.MustRegister(responseSizeHistogram)
}

func RegisterMetrics() error {
	return view.Register(
		&view.View{
			Name:        "service/content/total_requests",
			Description: "Count of HTTP requests received",
			TagKeys:     []tag.Key{ProjectIDKey, HostKey, EndpointKey, ClientIPKey},
			Measure:     TotalRequests,
			Aggregation: view.Count(),
		},
		&view.View{
			Name:        "service/content/latency_distribution",
			Description: "Latency distribution of HTTP requests",
			Measure:     RequestLatency,
			TagKeys:     []tag.Key{EndpointKey, StatusKey, ClientIPKey},
			Aggregation: view.Distribution(0, 50, 100, 200, 500, 1000, 2000),
		},
		&view.View{
			Name:        "service/content/request_size_distribution",
			Description: "Distribution of HTTP request sizes",
			Measure:     RequestSizeBytes,
			TagKeys:     []tag.Key{MethodKey, ClientIPKey},
			Aggregation: view.Distribution(100, 200, 500, 1000, 5000),
		},
		&view.View{
			Name:        "service/content/response_size_distribution",
			Description: "Distribution of HTTP response sizes",
			Measure:     ResponseSizeBytes,
			TagKeys:     []tag.Key{StatusKey, EndpointKey, ClientIPKey},
			Aggregation: view.Distribution(100, 200, 500, 1000, 5000),
		},
	)
}

func RecordContentRequest(projectID string, host string, clientIP net.IP, endpoint string, status int, latency time.Duration) {
	hostname := os.Getenv("NODE_ID")
	if hostname == "" {
		hostname = "unknown"
	}

	clientIPStr := clientIP.String()

	// Record total requests with all labels
	totalRequests.WithLabelValues(
		hostname,
		clientIPStr,
		projectID,
		endpoint,
		host,
	).Inc()

	// Record latency with additional labels
	latencyHistogram.WithLabelValues(
		hostname,
		clientIPStr,
		endpoint,
	).Observe(latency.Seconds())

	// Record request and response size distributions with additional labels
	requestSizeHistogram.WithLabelValues(
		hostname,
		clientIPStr,
		endpoint,
	).Observe(float64(0))

	responseSizeHistogram.WithLabelValues(
		hostname,
		clientIPStr,
		endpoint,
	).Observe(float64(0))

	// Create context with tags for OpenCensus metrics
	ctx, _ := tag.New(context.Background(),
		tag.Insert(ProjectIDKey, projectID),
		tag.Insert(HostKey, host),
		tag.Insert(EndpointKey, endpoint),
		tag.Insert(StatusKey, fmt.Sprintf("%d", status)),
		tag.Insert(ClientIPKey, clientIPStr),
	)

	// Record OpenCensus metrics
	stats.Record(ctx,
		TotalRequests.M(1),
		RequestLatency.M(float64(latency.Milliseconds())),
	)
}

func StartMetricsServer(logger *zap.Logger, port string) error {
	if err := RegisterMetrics(); err != nil {
		return fmt.Errorf("failed to register metrics: %w", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	logger.Info("Starting metrics server", zap.String("port", port))
	return http.ListenAndServe(":"+port, nil)
}

func GenerateTestMetrics(logger *zap.Logger, duration time.Duration) {
	logger.Info("Starting test metrics generation", zap.Duration("duration", duration))

	projects := []string{"project_1", "project_2", "project_3"}
	hosts := []string{
		"cdn1.example.com",
		"cdn2.example.com",
		"cdn3.example.com",
	}

	generateBatch := func() {
		for _, project := range projects {
			for _, host := range hosts {
				clientIP := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
				logger.Debug("Recording metric",
					zap.String("endpoint", "/api/content/test"),
					zap.String("host", host),
					zap.String("project_id", project),
					zap.String("client_ip", clientIP),
				)
				RecordContentRequest(
					project,
					host,
					net.ParseIP(clientIP),
					"/api/content/test",
					200,
					time.Duration(50+rand.ExpFloat64()*150)*time.Millisecond,
				)
			}
		}
	}

	// Generate initial batch
	generateBatch()
	logger.Info("Initial batch of metrics generated")

	if duration > 0 {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		timeout := time.After(duration)
		count := 0

		for {
			select {
			case <-ticker.C:
				generateBatch()
				count++
				if count%100 == 0 {
					logger.Info("Generated metrics batches", zap.Int("count", count))
				}
			case <-timeout:
				logger.Info("Finished generating test metrics", zap.Int("total_batches", count))
				return
			}
		}
	}
}
