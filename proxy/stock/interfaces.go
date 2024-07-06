package stock

import "context"

type StockService interface {
	StockListsService(ctx context.Context) (ApiResponse, error)
	StockPredictionService(ctx context.Context, req PredictionRequest) (ApiResponse, error)
}

type StockServiceImpl struct{}