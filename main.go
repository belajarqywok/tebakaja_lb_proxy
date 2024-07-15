package main

import (
	"os"
	"fmt"
	"log"
	
	"github.com/joho/godotenv"

	// Fiber
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"

	// Main Features
	stock_proxy             "tebakaja_lb_proxy/proxy/stock"
	crypto_proxy            "tebakaja_lb_proxy/proxy/crypto"
	national_currency_proxy "tebakaja_lb_proxy/proxy/national_currency"

	middlewares             "tebakaja_lb_proxy/proxy/middlewares"

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
	err := godotenv.Load()
  if err != nil {
    log.Fatalf("Error loading .env file")
  }

	proxyService := fiber.New()
	proxyService.Use(helmet.New())
	proxyService.Use(middlewares.LoggingMiddleware)
	proxyService.Use(middlewares.RateLimiterMiddleware())

	proxyService.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("TEBAKAJA_CORS_ALLOW_ORIGINS"),
		AllowHeaders: os.Getenv("TEBAKAJA_CORS_ALLOW_HEADERS"),
		AllowMethods: os.Getenv("TEBAKAJA_CORS_ALLOW_METHODS"),
		AllowCredentials: true,
	}))

	proxyService.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Security-Policy", fmt.Sprintf("frame-ancestors 'self' %s %s %s %s %s %s",
			"https://huggingface.co",
			"https://qywok-tebakaja-proxy-space-0.hf.space",
			"https://qywok-tebakaja-proxy-space-1.hf.space",
			"https://qywok-tebakaja-proxy-space-2.hf.space",
			"https://qywok-tebakaja-proxy-space-3.hf.space",
			"https://qywok-tebakaja-proxy-space-4.hf.space",
		))
		return c.Next()
	})

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

	HOST := os.Getenv("TEBAKAJA_PROXY_HOST")
	PORT := os.Getenv("TEBAKAJA_PROXY_PORT")
	log.Fatal(proxyService.Listen(fmt.Sprintf("%s:%s", HOST, PORT)))
}
