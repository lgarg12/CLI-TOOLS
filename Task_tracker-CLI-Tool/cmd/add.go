package cmd


import (
	"fmt"

	"github.com/spf13/cobra"
	"Task-Tracker-CLI-Tool/V1/store" // Replace with actual path to your task logic
)

var description string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  "Add a new task to your task list with a default status of 'todo'.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task name.")
			return
		}
		name := args[0] 
		store.AddTask(name, description)
		fmt.Println("Task added successfully.")
	},
}

func init() {
	addCmd.Flags().StringVarP(&description,"description","d","","add task description")
	rootCmd.AddCommand(addCmd)
}