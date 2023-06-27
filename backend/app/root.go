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

// RunServer godoc
//
//	@title			NutBooks API
//	@version		1.0.0
//	@contact.email	cheesecat47@gmail.com
//	@licence.name	MIT
func RunServer() {

	config := configs.FiberConfig()

	app := fiber.New(config)

	middlewares.FiberMiddleware(app)

	// limit 3 requests per 10 seconds max
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        10,
	}))

	// https://github.com/swaggo/swag#declarative-comments-format
	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)

	// Graceful shutdown
	// https://github.com/gofiber/recipes/tree/master/graceful-shutdown
	go func() {
		if err := app.Listen(":3000"); err != nil {
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
