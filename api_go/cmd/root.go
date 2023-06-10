package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

// ----------------------------------------------------------------
//
// Base
//
// ----------------------------------------------------------------

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	//rootCmd.PersistentFlags().StringP("author", "a", "cheesecat47", "Author name for copyright attribution")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.SetDefault("author", "cheesecat47 <cheesecat47@gmail.com>")

	rootCmd.AddCommand(runCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("err")
		os.Exit(1)
	}
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
		fmt.Println("" +
			" ____              _                         _        \n" +
			"| __ )  ___   ___ | | ___ __ ___   __ _ _ __| | _____ \n" +
			"|  _ \\ / _ \\ / _ \\| |/ / '_ ` _ \\ / _` | '__| |/ / __|\n" +
			"| |_) | (_) | (_) |   <| | | | | | (_| | |  |   <\\__ \\\n" +
			"|____/ \\___/ \\___/|_|\\_\\_| |_| |_|\\__,_|_|  |_|\\_\\___/")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run Server: " + strings.Join(args, " "))
	},
}
