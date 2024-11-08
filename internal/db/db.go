package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func Init() {
	db, err := sql.Open("sqlite", "todos.db")
	if err != nil {
		fmt.Println("Error creating database")
		os.Exit(1)
	}
	defer db.Close()

	// create table if not exists
	createTodosTable := `CREATE TABLE IF NOT EXISTS todos (
    id STRING PRIMARY KEY,
    title STRING NOT NULL,
    status STRING NOT NULL,
    createdat DATETIME,
    updatedat DATETIME
  )
  `

	_, err = db.Exec(createTodosTable)

	if err != nil {
		fmt.Println("Error while creating the table")
		os.Exit(1)
	}

}
