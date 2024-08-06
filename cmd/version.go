package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of trueselfe-cli",
	Long:  `All software has versions. This is the True Selfe CLI's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("True Selfe CLI v0.1 -- HEAD")
	},
}
