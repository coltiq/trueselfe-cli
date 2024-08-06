/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/coltiq/trueselfe-cli/cmd"
	"github.com/coltiq/trueselfe-cli/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
