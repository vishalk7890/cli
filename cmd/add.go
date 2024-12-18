package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"additions"},
	Long:    "doing addition via maths",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Addition of %s and %s = %s.\n\n", args[0], args[1], Add(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
