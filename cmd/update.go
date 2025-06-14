package cmd

import (
	"fmt"
    "strings"
	"Task-Tracker-CLI-Tool/V1/store"

	"github.com/spf13/cobra"
)

var (
	nameFlag        string
	descriptionFlag string
	statusFlag      string
)

var updateCmd = &cobra.Command{
	Use:   "update <taskId>",
	Short: "Update a task",
	Long: `Update a task by its ID. 
You can selectively update the name, description, and/or status using flags.
Valid statuses: todo, in-progress, done.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("❌ Please provide a task ID.")
			return
		}

		taskId := args[0]

		if statusFlag != "" {
			statusFlag = strings.ToLower(statusFlag)
			if statusFlag != "todo" && statusFlag != "in-progress" && statusFlag != "done" {
				fmt.Println("❌ Invalid status. Allowed values: todo, in-progress, done")
				return
			}
		}

		// Call your update logic
		store.UpdateTask(taskId, nameFlag, descriptionFlag, statusFlag)

		fmt.Println("✅ Task updated successfully.")
	},
}

func init() {
	updateCmd.Flags().StringVarP(&nameFlag, "name", "n", "", "Update task name")
	updateCmd.Flags().StringVarP(&descriptionFlag, "description", "d", "", "Update task description")
	updateCmd.Flags().StringVarP(&statusFlag, "status", "s", "", "Update task status (todo, in-progress, done)")
	rootCmd.AddCommand(updateCmd)
}
