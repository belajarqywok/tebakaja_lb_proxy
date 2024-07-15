package crypto

import (
	"fmt"
	"bytes"
	"context"
	"net/http"
	"encoding/json"

	helpers "tebakaja_lb_proxy/proxy/helpers"
)


/*
 * --- Cryptocurrency Prediction Service ---
 */
func (s *CryptoServiceImpl) CryptoPredictionService(ctx context.Context, req PredictionRequest) (ApiResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return ApiResponse{
			Message:    fmt.Sprintf("Failed to marshal request body: %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	endpoint := fmt.Sprintf("%s/prediction", helpers.GetEndpointService("crypto"))
	httpReq, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return ApiResponse{
			Message:    fmt.Sprintf("Failed to make request: %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return ApiResponse{
			Message:    fmt.Sprintf("Failed to make request: %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ApiResponse{
			Message:    fmt.Sprintf("Request failed: %d", resp.StatusCode),
			StatusCode: resp.StatusCode,
		}, fmt.Errorf("request failed: %d", resp.StatusCode)
	}

	var apiResponse PredictionResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return ApiResponse{
			Message:    fmt.Sprintf("Failed to parse JSON response: %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return ApiResponse{
		Message:    apiResponse.Message,
		StatusCode: apiResponse.StatusCode,
		Data:       apiResponse.Data,
	}, nil
}
