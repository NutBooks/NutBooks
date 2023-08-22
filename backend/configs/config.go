package configs

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
)

var (
	JWTSecret string
)

func Config() {
	var exists bool
	JWTSecret, exists = os.LookupEnv("JWT_SECRET")
	if !exists {
		log.Fatalw("[func Protected] Missing JWT Secret", "JWTSecret", JWTSecret)
		os.Exit(1)
	}
}
