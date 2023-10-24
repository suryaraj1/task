package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		task, err := db.DeleteTask(id)

		if err != nil {
			fmt.Println("Encountered error when trying to delete task from db!!")
			return
		}

		fmt.Println(task.Value + ": Deletion success!")

	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
