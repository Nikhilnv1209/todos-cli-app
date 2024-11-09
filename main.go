package main

import (
	"todos-app/cmd"
	"todos-app/internal/db"
)

func main() {
	db.Init()

	defer db.Close()

	cmd.Execute()
}
