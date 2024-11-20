package cmd

import (
	"fmt"
	"todos-app/internal/db"

	"github.com/spf13/cobra"
)

var deleteId, deleteTitle, deleteStatus string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete any specific or all the todos in the todos list",
	Run: func(cmd *cobra.Command, args []string) {
		if deleteId == "" && deleteStatus == "" && deleteTitle == "" {
			fmt.Println("Atleast specify one flag from Id, title, and status to delete a todo")
			return
		}

		delete(deleteId, deleteTitle, deleteStatus)
	},
}

func delete(id string, title string, status string) {
	deleteQuery := `DELETE FROM todos WHERE 1=1`
	args := []interface{}{}

	if id != "" {
		deleteQuery += " AND id = ?"
		args = append(args, id)
	}

	if title != "" {
		deleteQuery += " AND title = ?"
		args = append(args, title)
	}

	if status != "" {
		deleteQuery += " AND status = ?"
		args = append(args, status)
	}

	fmt.Println(deleteQuery, args)

	result, err := db.Client.Exec(deleteQuery, args...)

	if err != nil {
		fmt.Printf("error deleting rows %v", err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected > 0 {
		fmt.Printf("Deleted %d Todos succesfully", rowsAffected)
	} else {
		fmt.Println("No todos match the criteria")
	}

}

func init() {
	deleteCmd.Flags().StringVar(&deleteId, "id", "", "Provide the id of the todo that you want to delete")
	deleteCmd.Flags().StringVar(&deleteTitle, "title", "", "Provide the title of the todo that you want to delete")
	deleteCmd.Flags().StringVar(&deleteStatus, "status", "", "Provide the status of the todo that you want to delete")
}
