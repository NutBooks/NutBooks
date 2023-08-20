package middlewares

import (
	"api/db/models"
	"errors"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var (
	JWTHandler *fiber.Handler
	JWTSecret  string
)

func Protected() fiber.Handler {
	if JWTHandler == nil {
		JWTSecret, exists := os.LookupEnv("JWT_SECRET")
		if !exists {
			log.Fatalw("[func Protected] Missing JWT Secret", "JWTSecret", JWTSecret)
			os.Exit(1)
		}
		newJwtHandler := jwtware.New(jwtware.Config{
			SigningKey:   jwtware.SigningKey{Key: []byte(JWTSecret)},
			ErrorHandler: jwtErrorHandler,
		})
		JWTHandler = &newJwtHandler
	}
	return *JWTHandler
}

func ValidToken(t *jwt.Token, id uint) error {
	claims := t.Claims.(jwt.MapClaims)
	if claims["user_id"] == id {
		return nil
	} else {
		return errors.New("Invalid token")
	}
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
