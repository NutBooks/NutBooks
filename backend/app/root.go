// ----------------------------------------------------------------
//
// API Server
// https://docs.gofiber.io/
//
// ----------------------------------------------------------------

package app

import (
	"api/app/middlewares"
	"api/configs"
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

// RunServer godoc
// @title			NutBooks API
// @version			1.0.0
// @contact.email	cheesecat47@gmail.com
// @licence.name	MIT
func RunServer() {

	config := configs.FiberConfig()

	app := fiber.New(config)

	middlewares.FiberMiddleware(app)

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

// ApiRoot godoc
// @Summary	Root URL - for health check
// @Success	200
// @Router		/ [get]
func ApiRoot(c *fiber.Ctx) error {
	return c.SendString("Hi")
}
