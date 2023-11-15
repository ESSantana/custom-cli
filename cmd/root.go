package cmd

import (
	"os"

	"github.com/ESSantana/custom-cli/cmd/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "custom-cli",
	Short: "Minha primeira CLI custom",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	commands.InitAll(rootCmd)
}
