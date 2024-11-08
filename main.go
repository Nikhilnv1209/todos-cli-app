package main

import (
	"todos-app/cmd"
	"todos-app/internal/db"
)

func main() {
	db.Init()
	cmd.Execute()
}
