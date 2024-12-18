package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:     "sub",
	Aliases: []string{"additions"},
	Long:    "doing subration via maths",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Subraction of %s and %s = %s.\n\n", args[0], args[1], Subtract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
