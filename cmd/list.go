package cmd

import (
	"fmt"
	"todos-app/internal/db"
	"todos-app/internal/types"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List down all the todos in the form of a table",
	Run: func(cmd *cobra.Command, args []string) {
		List()
	},
}

func List() {
	listQuery := `SELECT * FROM todos`

	result, err := db.Client.Query(listQuery)

	if err != nil {
		fmt.Printf("error while listing the todos %v", err)
		return
	}

	defer result.Close()

	hasRows := false

	for result.Next() {
		hasRows = true
		var todo types.Todo
		err := result.Scan(&todo.Id, &todo.Title, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)

		if err != nil {
			fmt.Println("error while scanning the row")
			continue
		}

		fmt.Println(todo.Id, todo.Title, todo.Status, todo.CreatedAt, todo.UpdatedAt)
	}

	if !hasRows {
		fmt.Println("No rows found")
	}

	if err = result.Err(); err != nil {
		fmt.Println("Error reading rows", err)
	}
}
