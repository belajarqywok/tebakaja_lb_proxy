package crypto

type ApiResponse struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}


type PredictionRequest struct {
	Days     int    `json:"days" validate:"gte=1,lte=31"`
	Currency string `json:"currency" validate:"required,min=4,max=16"`
}


type PredictionResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`

	Data       struct {
		Currency    string `json:"currency"`
		Predictions struct {
			Actuals    []struct {
				Date  string  `json:"date"`
				Price float64 `json:"price"`
			} `json:"actuals"`
			Predictions []struct {
				Date  string  `json:"date"`
				Price float64 `json:"price"`
			} `json:"predictions"`
		} `json:"predictions"`
	} `json:"data"`
}
