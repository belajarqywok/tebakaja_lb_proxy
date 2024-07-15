package middlewares

import (
	"log"
	"time"
	"net/http"

	"github.com/gofiber/fiber/v2"
)


/*
 *  --- Logging Middleware ---
*/
func LoggingMiddleware(c *fiber.Ctx) error {
	start_time := time.Now()

	log.Printf("[%s] %s %s - %d %s in %v",
		time.Now().Format("2006-01-02 15:04:05"),
			c.Method(), c.Path(), c.Response().StatusCode(),
				http.StatusText(c.Response().StatusCode()), time.Since(start_time))

				
	return c.Next()
}

