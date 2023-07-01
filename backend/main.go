package main

import (
	"api/app"
	"api/db"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
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
		app.RunServer()
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
