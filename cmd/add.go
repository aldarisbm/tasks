package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to your tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		taskToAdd := strings.Join(args, " ")
		fmt.Printf("Added to your tasks list: \"%s\"", taskToAdd)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
