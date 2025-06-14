package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"Task-Tracker-CLI-Tool/V1/store"
)

var status string

var printTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "Print Task List",
	Long:  "List all tasks or filter them by status (todo, in-progress, done).",
	Run: func(cmd *cobra.Command, args []string) {
		if status == "" {
			store.PrintAllTasks()
		} else {
			filteredStatus := strings.ToLower(status)
			if filteredStatus != "todo" && filteredStatus != "in-progress" && filteredStatus != "done" {
				fmt.Println("❌ Invalid status. Please use: todo, in-progress, or done.")
				return
			}
			store.PrintTasksByStatus(filteredStatus)
		}
	},
}

func init() {
	rootCmd.AddCommand(printTasksCmd)
	printTasksCmd.Flags().StringVarP(&status, "status", "s", "", "Filter tasks by status (todo, in-progress, done)")
}
