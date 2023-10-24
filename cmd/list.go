package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Encountered error while trying to read tasks from db!!")
			return
		}

		fmt.Println("You have the following tasks:")
		for id, task := range tasks {
			fmt.Println(strconv.Itoa(id+1) + ". " + task.Value + " ID: " + strconv.Itoa(task.Key))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
