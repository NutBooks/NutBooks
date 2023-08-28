package middlewares

import (
	"api/configs"
	"api/db/models"
	"errors"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(configs.JWTSecret)},
		ErrorHandler: jwtErrorHandler,
	})
}

func ValidToken(t *jwt.Token, email string) error {
	claims := t.Claims.(jwt.MapClaims)
	if claims["email"] == email {
		return nil
	} else {
		return errors.New("invalid token")
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
