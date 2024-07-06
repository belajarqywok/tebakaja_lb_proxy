package proxy

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// Logging Middleware
func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err   := c.Next()

	log.Printf("[%s] %s %s - %d %s in %v", time.Now().Format("2006-01-02 15:04:05"),
		c.Method(), c.Path(), c.Response().StatusCode(),
		http.StatusText(c.Response().StatusCode()), time.Since(start))

	return err
}


/*
 *  --- Rate Limiter Middleware ---
*/
func RateLimiterMiddleware() func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        200,
		Expiration: 1 * time.Minute,

		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},

		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded",
			})
		},
	})
}
