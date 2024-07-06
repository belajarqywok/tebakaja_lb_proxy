package national_currency

import "context"

type NationalCurrencyService interface {
	NationalCurrencyListsService(ctx context.Context) (ApiResponse, error)
	NationalCurrencyPredictionService(ctx context.Context, req PredictionRequest) (ApiResponse, error)
}

type NationalCurrencyServiceImpl struct{}