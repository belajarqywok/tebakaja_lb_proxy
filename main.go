package main

import (
	"fmt"
	"log"
	
	"github.com/gofiber/fiber/v2"

	proxy "tebakaja_lb_proxy/proxy"
	crypto_proxy "tebakaja_lb_proxy/proxy/crypto"
)

func main() {
	app := fiber.New()

	app.Use(proxy.LoggingMiddleware)
	app.Use(proxy.RateLimiter())

	cryptoGroup := app.Group("/crypto")
	cryptoGroup.Get("/lists", crypto_proxy.CryptoListsHandler(&crypto_proxy.CryptoServiceImpl{}))
	cryptoGroup.Post("/prediction", crypto_proxy.CryptoPredictionHandler(&crypto_proxy.CryptoServiceImpl{}))

	port := 7860
	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%d", port)))
}
