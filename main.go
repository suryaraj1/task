package main

import (
	"path/filepath"
	"task/cmd"
	"task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {

	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)

	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
