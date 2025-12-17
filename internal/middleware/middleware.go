package middleware

import (
	"user-api/internal/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(requestid.New())

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		logger.Log.Info("HTTP Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.String("request_id", c.GetRespHeader("X-Request-ID")),
			zap.Duration("duration", duration),
		)

		return err
	})
}