package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This will add the todos in the main todo list",
}

func init() {
	todoCmd.AddCommand(addCmd)
}
