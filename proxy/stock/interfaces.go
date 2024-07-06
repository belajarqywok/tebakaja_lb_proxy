package stock

import "context"

type CryptoService interface {
	CryptoListsService(ctx context.Context) (ApiResponse, error)
	CryptoPredictionService(ctx context.Context, req PredictionRequest) (ApiResponse, error)
}

type CryptoServiceImpl struct{}