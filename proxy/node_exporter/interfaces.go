package exporter

import "context"

type ExporterService interface {
	ExporterMetricsService(ctx context.Context) (string, error)
	// ExporterVersionInfoService(ctx context.Context) (string, error)
	// ExporterHealthCheckService(ctx context.Context) (string, error)
}

type ExporterServiceImpl struct{}