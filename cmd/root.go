/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example/com/app/service"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var option string
var opt string

var drawCmd = &cobra.Command{

	Use:     "draw",
	Aliases: []string{"ref"},
	Short:   "Refreshes Current Coin Details",
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		service.CreateTable(service.CreateData(service.FindCoinFromIndexMap(args[0])))

		if option == "refresh" {

			time.Sleep(8 * time.Second)

			service.CreateTable(service.CreateData(service.FindCoinFromIndexMap(args[0])))

		}

		if option == "change" {

			time.Sleep(8 * time.Second)

			var coinName string

			fmt.Scanf("%s", &coinName)

			coinName = strings.ToUpper(coinName)

			service.CreateTable(service.CreateData(service.FindCoinFromIndexMap(coinName)))
		}

		if option == "exit" {

			log.Fatal("Exiting")

		}

		//log.Fatal("Undefined choice, Exiting...")

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(drawCmd)
	//drawCmd.Flags().StringVarP(&option, "options", "t", "options for outputs", "idk")
	drawCmd.PersistentFlags().StringVar(&option, "options", "o", "options outputs")

}
