package cmd

import (
	"fmt"
	"github.com/aldarisbm/tasks/db"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to your tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		taskToAdd := strings.Join(args, " ")
		_, err := db.CreateTask(taskToAdd)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		fmt.Printf("Added to your tasks list: \"%s\"", taskToAdd)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
