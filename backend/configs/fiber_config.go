package configs

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		IdleTimeout: 5 * time.Second,
		ReadTimeout: 60 * time.Second,
	}
}
