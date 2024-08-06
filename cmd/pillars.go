package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pillarsCmd = &cobra.Command{
	Use:   "pillars",
	Short: "Physical Health, Mental Health, Personal Growth, Relationships, Professional Development",
	Long: `Welcome to True Selfe! To help you achieve a well-rounded and fulfilling life, we've identified five key pillars of self-improvement: 
Physical Health, Mental Health, Personal Growth, Relationships, Professional Development`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`1. Physical Health: Prioritize your body's well-being through exercise, nutrition, and proper rest.
2. Mental Health: Foster a positive mindset with mindfulness, stress management, and emotional resilience.
3. Personal Growth: Expand your horizons by learning new skills, pursuing hobbies, and engaging in continuous education.
4. Relationships: Strengthen your connections with family, friends, and colleagues through effective communication and meaningful interactions.
5. Professional Development: Advance your career and enhance productivity with goal setting, efficient time management, and work-life balance strategies.`)
	},
}

func init() {
	rootCmd.AddCommand(pillarsCmd)
}
