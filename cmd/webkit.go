/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/tools_scraping_py/internal/controller"
)

// webkitCmd represents the webkit command
var webkitCmd = &cobra.Command{
	Use:   "webkit",
	Short: "Select your web api",
	Run: func(cmd *cobra.Command, args []string) {
		controller.WebApi()
	},
}

func init() {
	rootCmd.AddCommand(webkitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webkitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webkitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
