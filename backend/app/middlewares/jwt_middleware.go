package middlewares

import (
	"api/db/models"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
)

var (
	JWTHandler *fiber.Handler
)

func Protected() fiber.Handler {
	if JWTHandler == nil {
		jwtSecret, exists := os.LookupEnv("JWT_SECRET")
		if !exists {
			log.Fatalw("[func Protected] Missing JWT Secret", "jwtSecret", jwtSecret)
			os.Exit(1)
		}
		newJwtHandler := jwtware.New(jwtware.Config{
			SigningKey:   jwtware.SigningKey{Key: []byte(jwtSecret)},
			ErrorHandler: jwtErrorHandler,
		})
		JWTHandler = &newJwtHandler
	}
	return *JWTHandler
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(models.JwtErrorResponse{
			Error:   true,
			Message: "Missing or malformed JWT",
			Data:    err,
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(models.JwtErrorResponse{
		Error:   true,
		Message: "Invalid or expired JWT",
		Data:    err,
	})
}
