package cmd

import (
	"fmt"
	"time"
	"todos-app/internal/db"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var addTitle, addStatus string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This will add the todos in the main todo list",
	Run: func(cmd *cobra.Command, args []string) {
		if addTitle == "" {
			fmt.Println("please specify the title for the todos")
			return
		}

		status, err := addTodos(addTitle, addStatus)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(status)
		}
	},
}

func addTodos(title string, status string) (string, error) {
	id := uuid.NewString()
	createdAt := time.Now().Format("2006-01-02 03:04:05 PM")
	updatedAt := createdAt

	insertQuery := `INSERT INTO todos (id, title, status, createdAt, updatedAt) VALUES (?,?,?,?,?)`

	_, err := db.Client.Exec(insertQuery, id, title, status, createdAt, updatedAt)
	if err != nil {
		return "", fmt.Errorf("error adding todo %v", err)
	}

	return "todo added succesfully", nil
}

func init() {
	addCmd.Flags().StringVarP(&addTitle, "title", "t", "", "Title of the todo")
	addCmd.Flags().StringVarP(&addStatus, "status", "s", "Pending", "Status for the todos")
}
