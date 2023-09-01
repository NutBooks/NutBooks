package main

import (
	"api/app/middlewares"
	"api/app/routes"
	"api/configs"
	"api/db"
	_ "api/docs"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("err")
		os.Exit(1)
	}
}

// ----------------------------------------------------------------
//
// Base
// https://cobra.dev/#getting-started
//
// ----------------------------------------------------------------

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	//rootCmd.PersistentFlags().StringP("author", "a", "cheesecat47", "Author name for copyright attribution")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.SetDefault("author", "cheesecat47 <cheesecat47@gmail.com>")

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(migrateDB)
}

// ----------------------------------------------------------------
//
// Commands
//
// ----------------------------------------------------------------

var runCmd = &cobra.Command{
	Use:   "run [options]",
	Short: "Run server",
	Long:  `Run an api server`,
	Args:  cobra.MinimumNArgs(0),
	PersistentPreRun: func(c *cobra.Command, args []string) {
		fmt.Println(
			"    _   __      __  ____              __\n" +
				"   / | / /_  __/ /_/ __ )____  ____  / /_______\n" +
				"  /  |/ / / / / __/ __  / __ \\/ __ \\/ //_/ ___/\n" +
				" / /|  / /_/ / /_/ /_/ / /_/ / /_/ / ,< (__  )\n" +
				"/_/ |_/\\__,_/\\__/_____/\\____/\\____/_/|_/____/",
		)
	},
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

var migrateDB = &cobra.Command{
	Use:   "migrate [options]",
	Short: "Migrate DB",
	Long:  `Migrate MySQL database`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Migrate")
		db.MigrateMysql()
	},
}

// ----------------------------------------------------------------
//
// API Server
// https://docs.gofiber.io/
//
// ----------------------------------------------------------------

//	@title			NutBooks API
//	@version		1.0.0
//	@description	Nutbooks API documentation

//	@contact.email	cheesecat47@gmail.com
//	@licence.name	MIT

//	@BasePath	/api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				AccessToken
func runServer() {
	configs.Config()

	db.Connect()

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
		ReadTimeout: 60 * time.Second,
	})

	middlewares.CorsMiddleware(app)
	middlewares.FiberMiddleware(app)

	// limit 3 requests per 10 seconds max
	app.Use(limiter.New(limiter.Config{
		Expiration: 1 * time.Hour,
		Max:        1000,
	}))

	// https://github.com/swaggo/swag#declarative-comments-format
	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)

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
