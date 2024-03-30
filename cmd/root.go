/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tools_scraping_py",
	Short: "A toolkit for you build-in your apps in python",
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

	fmt.Println(" _____           _       ____                       _               ____        ")
	fmt.Println("|_   _|__   ___ | |___  / ___|  ___ _ __ __ _ _ __ (_)_ __   __ _  |  _ \\ _   _ ")
	fmt.Println("  | |/ _ \\ / _ \\| / __| \\___ \\ / __| '__/ _` | '_ \\| | '_ \\ / _` | | |_) | | | |")
	fmt.Println("  | | (_) | (_) | \\__ \\  ___) | (__| | | (_| | |_) | | | | | (_| | |  __/| |_| |")
	fmt.Println("  |_|\\___/ \\___/|_|___/ |____/ \\___|_|  \\__,_| .__/|_|_| |_|\\__, | |_|    \\__, |")
	fmt.Println("                                             |_|            |___/         |___/ ")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tools_scraping_py.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
