package main

import (
	"fmt"
	"log"
	
	"github.com/gofiber/fiber/v2"

	proxy "tebakaja_lb_proxy/proxy"

	// Main Features
	stock_proxy "tebakaja_lb_proxy/proxy/stock"
	crypto_proxy "tebakaja_lb_proxy/proxy/crypto"
	national_currency_proxy "tebakaja_lb_proxy/proxy/national_currency"

	// Swagger
	_ "tebakaja_lb_proxy/docs"
	swagger "github.com/swaggo/fiber-swagger"

	// Node Exporter
	// exporter_proxy "tebakaja_lb_proxy/proxy/node_exporter"
)


// @title          TebakAja
// @version        1.0
// @description    TebakAja REST API Service
// @termsOfService https://swagger.io/terms/

// @contact.name   Si Mimin
// @contact.url    https://www.tebakaja.com
// @contact.email  tebakaja@gmail.com

// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.137.1:7860
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

	proxyService.Get("/swagger/*", swagger.WrapHandler)
	proxyService.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", fiber.StatusMovedPermanently)
	})

	host := "0.0.0.0"
	port := 7860
	log.Fatal(proxyService.Listen(fmt.Sprintf("%s:%d", host, port)))
}
