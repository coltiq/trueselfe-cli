/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "trueselfe-cli",
		Short: "Tracker for your own self improvement journey",
		Long: `True Selfe CLI is a powerful command-line tool designed to help you achieve holistic self-improvement by tracking your 
progress across five key pillars of personal development. This user-friendly application allows you to set, manage, and monitor your goals, 
ensuring that you stay on track and make meaningful progress in all aspects of your life.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Name() != "init" && !isInitDone() {
				fmt.Println("Error: 'init' command must be run before any other commands.")
				os.Exit(1)
			}
		},
	}
)

func isInitDone() bool {
	_, err := os.Stat(".initialized")
	return !errors.Is(err, os.ErrNotExist)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
