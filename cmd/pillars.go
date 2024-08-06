package cmd

import (
	"fmt"

	"github.com/coltiq/trueselfe-cli/data"
	"github.com/spf13/cobra"
)

var pillarsCmd = &cobra.Command{
	Use:   "pillars",
	Short: "Physical Health, Mental Health, Personal Growth, Relationships, Professional Development",
	Long: `Welcome to True Selfe! To help you achieve a well-rounded and fulfilling life, we've identified five key pillars of self-improvement: 
Physical Health, Mental Health, Personal Growth, Relationships, Professional Development`,
	Run: func(cmd *cobra.Command, args []string) {
		pillars := data.GetPillars()
		fmt.Println()
		for i, pillar := range pillars {
			fmt.Printf("%d. %s: %s\n", i+1, pillar.Name, pillar.Description)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(pillarsCmd)
}
