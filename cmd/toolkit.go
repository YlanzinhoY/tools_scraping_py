/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/tools_scraping_py/internal/controller"
)

// toolkitCmd represents the toolkit command
var toolkitCmd = &cobra.Command{
	Use:   "toolkit",
	Short: "Select tools for your web scraping",
	Run: func(cmd *cobra.Command, args []string) {
		controller.ToolkitController()
	},
}

func init() {
	rootCmd.AddCommand(toolkitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toolkitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toolkitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
