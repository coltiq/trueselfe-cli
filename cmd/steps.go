/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stepsCmd = &cobra.Command{
	Use:   "steps",
	Short: "Display all pillars and their associated steps",
	Long: `This allows you to display an overview of all your pillars along with the steps you created
to improve these pillars of your life.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("steps called")
	},
}

func init() {
	rootCmd.AddCommand(stepsCmd)
}
