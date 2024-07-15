package middlewares

import (
	"time"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)


/*
 *  --- Rate Limiter Middleware ---
*/
func RateLimiterMiddleware() func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        200,
		Expiration: (1 * time.Minute),

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

