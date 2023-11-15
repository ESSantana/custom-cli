package commands

import (
	"github.com/ESSantana/custom-cli/cmd/commands/calc"
	"github.com/spf13/cobra"
)

func InitAll(root *cobra.Command) {
	root.AddCommand(calc.Command())
}
