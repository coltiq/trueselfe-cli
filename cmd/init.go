/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/coltiq/trueselfe-cli/data"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new True Selfe database",
	Long:  `Initialize a new True Selfe database with tables: pillars, goals, completions.`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreatePillarTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
