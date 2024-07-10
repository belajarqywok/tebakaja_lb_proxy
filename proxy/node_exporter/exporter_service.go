package exporter

import (
	"context"
	"fmt"
	"io"
	"net/http"
)


/*
 *  --- Node Exporter Metric Service ---
 */
func (s *ExporterServiceImpl) ExporterMetricsService(ctx context.Context) (string, error) {
	endpoint := fmt.Sprintf("%s/metrics", "http://localhost:9100")
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	metricsResponse := string(body)

	return metricsResponse, nil
}
