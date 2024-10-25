package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "Todos",
	Short: "Todos is a cli app build using golang to provide a portable todos list in your system",
}

func init() {
	todoCmd.AddCommand(addCmd)
	todoCmd.AddCommand(listCmd)
	todoCmd.AddCommand(deleteCmd)
	todoCmd.AddCommand(updateCmd)
}

func Execute() {
	if err := todoCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
