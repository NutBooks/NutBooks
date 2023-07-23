// ----------------------------------------------------------------
//
// API Server
// https://docs.gofiber.io/
//
// ----------------------------------------------------------------

package app

import (
	"api/app/middlewares"
	"api/app/routes"
	"api/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "api/docs"
)

func Setup() *fiber.App {
	config := configs.FiberConfig()
	app := fiber.New(config)
	middlewares.FiberMiddleware(app)

	// limit 3 requests per 10 seconds max
	app.Use(limiter.New(limiter.Config{
		Expiration: 1 * time.Hour,
		Max:        1000,
	}))

	// https://github.com/swaggo/swag#declarative-comments-format
	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	return app
}

// RunServer godoc
//
//	@title			NutBooks API
//	@version		1.0.0
//	@contact.email	cheesecat47@gmail.com
//	@licence.name	MIT
func RunServer() {
	app := Setup()
	if app == nil {
		log.Panicf("Cannot create app")
	}

	// Graceful shutdown
	// https://github.com/gofiber/recipes/tree/master/graceful-shutdown
	go func() {
		port, exists := os.LookupEnv("API_PORT")
		if !exists {
			log.Println("No API_PORT environment variable. Use default value.")
			port = "8081"
		}
		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// notify at interrupt or termination
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	log.Printf("\nGraceful shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	// cleanup tasks
	// db.Close()
	log.Println("Fiber was successful shutdown.")
}
