package national_currency

import (
	"fmt"
	"log"
	"sync"
	"time"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)



// NationalCurrencyPredictionHandler  godoc
// @Summary                 National Currency Prediction
// @Description             National Currency Prediction
// @Tags                    National Currency
// @Accept                  json
// @Produce                 json
// @Param                   requestBody body PredictionRequest true "Request Body"
// @Success                 200  {object} PredictionResponse
// @Router                  /national-currency/lists [post]
func NationalCurrencyPredictionHandler(service NationalCurrencyService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 120*time.Second)
		defer cancel()

		ch := make(chan ApiResponse, 1)
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()

			var predictionReq PredictionRequest
			if err := c.BodyParser(&predictionReq); err != nil {
				log.Printf("[%s] Failed to parse request body: %v",
					time.Now().Format("2006-01-02 15:04:05"), err)

				ch <- ApiResponse{
					Message:    fmt.Sprintf("Failed to parse request body: %v", err),
					StatusCode: http.StatusBadRequest,
				}
				return
			}

			apiResponse, err := service.NationalCurrencyPredictionService(ctx, predictionReq)
			if err != nil {
				log.Printf("[%s] %v", time.Now().Format("2006-01-02 15:04:05"), err)
				ch <- apiResponse
				return
			}

			ch <- apiResponse
		}()

		go func() {
			wg.Wait()
			close(ch)
		}()

		select {
		case apiResponse, ok := <-ch:
			if !ok {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to get a response from the server",
				})
			}
			return c.Status(apiResponse.StatusCode).JSON(apiResponse)

		case <-ctx.Done():
			log.Printf("[%s] Timeout: %v", time.Now().Format("2006-01-02 15:04:05"), ctx.Err())
			return c.Status(http.StatusRequestTimeout).JSON(fiber.Map{
				"error": "Request timeout",
			})
		}
	}
}
