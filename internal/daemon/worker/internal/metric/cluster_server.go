package metric

import (
	"context"

	"github.com/hashicorp/boundary/globals"
	"github.com/hashicorp/boundary/internal/daemon/metric"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/stats"
)

const (
	workerClusterSubsystem = "worker_cluster"
)

// Because we use grpc's stats.Handler to track close-to-the-wire server-side communication over grpc,
// there is no easy way to measure request and response size as we are recording latency. Thus we only
// track the request latency for server-side grpc connections.

var grpcServerRequestLatency prometheus.ObserverVec = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: globals.MetricNamespace,
		Subsystem: workerClusterSubsystem,
		Name:      "grpc_request_duration_seconds",
		Help:      "Histogram of latencies for gRPC requests between the a worker server and a worker client.",
		Buckets:   prometheus.DefBuckets,
	},
	metric.ListGrpcLabels,
)

// All the codes expected to be returned by boundary or the grpc framework to
// requests to the cluster server.
var expectedGrpcCodes = []codes.Code{
	codes.OK, codes.InvalidArgument, codes.PermissionDenied,
	codes.FailedPrecondition,

	// Codes which can be generated by the gRPC framework
	codes.Canceled, codes.Unknown, codes.DeadlineExceeded,
	codes.ResourceExhausted, codes.Unimplemented, codes.Internal,
	codes.Unavailable, codes.Unauthenticated,
}

// InstrumentClusterStatsHandler returns a gRPC stats.Handler which observes
// cluster-specific metrics for a gRPC server.
func InstrumentClusterStatsHandler(ctx context.Context) (stats.Handler, error) {
	return metric.NewStatsHandler(ctx, grpcServerRequestLatency)
}

// InitializeClusterServerCollectors registers the cluster server metrics to the
// prometheus register and initializes them to 0 for all possible label
// combinations.
func InitializeClusterServerCollectors(r prometheus.Registerer, server *grpc.Server) {
	metric.InitializeGrpcCollectorsFromServer(r, grpcServerRequestLatency, server, expectedGrpcCodes)
}
