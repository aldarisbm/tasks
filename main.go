package main

import (
	"fmt"
	"github.com/aldarisbm/tasks/cmd"
	"github.com/aldarisbm/tasks/db"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := homedir.Dir()
	dbPath := filepath.Join(dir, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}