/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/coltiq/trueselfe-cli/data"
	"github.com/spf13/cobra"
)

var addParams []string

var stepsCmd = &cobra.Command{
	Use:   "steps",
	Short: "Display all pillars and their associated steps",
	Long: `This allows you to display an overview of all your pillars along with the steps you created
to improve these pillars of your life.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(addParams) == 4 {
			pillarID, err := strconv.Atoi(addParams[0])
			if err != nil {
				log.Fatal(err.Error())
			}
			name := addParams[1]
			description := addParams[2]
			frequency := addParams[3]
			data.AddStep(pillarID, name, description, frequency)
		} else if len(addParams) > 0 {
			log.Println("Error: -a flag requires exactly 4 parameters: pillarID, name, description, frequency")
		} else {
			fmt.Println("Steps Called")
		}
	},
}

func init() {
	rootCmd.AddCommand(stepsCmd)
	stepsCmd.Flags().StringSliceVarP(&addParams, "add", "a", nil, "Add a new step with parameters: pillarID, name, description, frequency")
}
