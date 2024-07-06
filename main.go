package main

import (
	"fmt"
	"log"
	
	"github.com/gofiber/fiber/v2"

	proxy "tebakaja_lb_proxy/proxy"

	stock_proxy "tebakaja_lb_proxy/proxy/stock"
	crypto_proxy "tebakaja_lb_proxy/proxy/crypto"
	national_currency_proxy "tebakaja_lb_proxy/proxy/national_currency"
)

func main() {
	proxyService := fiber.New()
	proxyService.Use(proxy.LoggingMiddleware)
	proxyService.Use(proxy.RateLimiterMiddleware())

	stockGroup := proxyService.Group("/stock")
	stockGroup.Get("/lists",
		stock_proxy.StockListsHandler(
			&stock_proxy.StockServiceImpl{}))
	stockGroup.Post("/prediction",
		stock_proxy.StockPredictionHandler(
			&stock_proxy.StockServiceImpl{}))

	cryptoGroup := proxyService.Group("/crypto")
	cryptoGroup.Get("/lists",
		crypto_proxy.CryptoListsHandler(
			&crypto_proxy.CryptoServiceImpl{}))
	cryptoGroup.Post("/prediction",
		crypto_proxy.CryptoPredictionHandler(
			&crypto_proxy.CryptoServiceImpl{}))

	nationalCurrencyGroup := proxyService.Group("/national-currency")
	nationalCurrencyGroup.Get("/lists",
		national_currency_proxy.NationalCurrencyListsHandler(
			&national_currency_proxy.NationalCurrencyServiceImpl{}))
	nationalCurrencyGroup.Post("/prediction",
		national_currency_proxy.NationalCurrencyPredictionHandler(
			&national_currency_proxy.NationalCurrencyServiceImpl{}))

	port := 7860
	log.Fatal(proxyService.Listen(fmt.Sprintf("0.0.0.0:%d", port)))
}
