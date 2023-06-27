// ----------------------------------------------------------------
//
// API Server
// https://docs.gofiber.io/
//
// ----------------------------------------------------------------

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "api/docs"
)

//	@title		Bookmarks API
//	@version	1.0.0

// @contact.email	cheesecat47@gmail.com
func RunServer() {

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	// limit 3 requests per 10 seconds max
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        10,
	}))

	// Swagger docs
	// https://github.com/gofiber/swagger
	// https://github.com/swaggo/swag
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/", ApiRoot)

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

// HealthCheck godoc
//
//	@Summary	Root URL - for health check
//	@Success	200
//	@Router		/ [get]
func ApiRoot(c *fiber.Ctx) error {
	return c.SendString("Hi")
}
