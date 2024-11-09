package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// Database connection
var (
	Client *sql.DB
	err    error
)

func Init() {
	Client, err = sql.Open("sqlite", "todos.db")
	if err != nil {
		fmt.Println("Error creating database")
		os.Exit(1)
	}

	// create table if not exists
	createTodosTable := `CREATE TABLE IF NOT EXISTS todos (
    id STRING PRIMARY KEY,
    title STRING NOT NULL,
    status STRING NOT NULL,
    createdat DATETIME,
    updatedat DATETIME
  )
  `

	_, err = Client.Exec(createTodosTable)

	if err != nil {
		fmt.Println("Error while creating the table")
		os.Exit(1)
	}

}

func Close() {
	if Client != nil {
		Client.Close()
	}
}
