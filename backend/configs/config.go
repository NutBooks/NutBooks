package configs

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
)

var (
	JWTSecret     string
	ALLOW_ORIGINS string
	FeDevHost     string
	CorsSecret    string
)

func Config() {
	var exists bool
	JWTSecret, exists = os.LookupEnv("JWT_SECRET")
	if !exists {
		log.Fatalw("[func Protected] Missing JWT Secret", "JWTSecret", JWTSecret)
		os.Exit(1)
	}

	ALLOW_ORIGINS, exists = os.LookupEnv("ALLOW_ORIGINS")
	if !exists {
		log.Fatalw("[func Protected] Missing ALLOW_ORIGINS", "ALLOW_ORIGINS", ALLOW_ORIGINS)
		os.Exit(1)
	}

	FeDevHost, exists = os.LookupEnv("FE_DEV_HOST")
	if !exists {
		FeDevHost = ""
	}

	CorsSecret, exists = os.LookupEnv("CORS_SECRET")
	if !exists {
		CorsSecret = ""
	}
}
