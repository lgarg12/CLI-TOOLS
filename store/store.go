package store

import (
	"Task-Tracker-CLI-Tool/V1/model"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

const fileName = "tasks.json"

var tasks []model.Task

func loadTasks() {
	tasks = []model.Task{} 

	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Failed to read file:", err)
		return
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
	}
}

func saveTasks() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal tasks:", err)
		return
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Failed to write file:", err)
	}
}

func AddTask(name string,description string) {
	loadTasks()
	now := time.Now().Format(time.RFC3339)

	newTask := model.Task{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Status:      model.StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, newTask)
	saveTasks()
}

func UpdateTask(id string, newName string, newDescription string, newStatus string) {
	loadTasks()

	updated := false
	for i, task := range tasks {
		if task.ID == id {
			if newName != "" {
                tasks[i].Name = newName
			}
			if newDescription != "" {
				tasks[i].Description = newDescription
			}
			if newStatus != "" {
				status := strings.ToLower(newStatus)
				if status == string(model.StatusTodo) ||
					status == string(model.StatusInProgress) ||
					status == string(model.StatusDone) {
					tasks[i].Status = model.TaskStatus(status)
				} else {
					fmt.Println("Invalid status. Must be one of: todo, in-progress, done.")
					return
				}
			}
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			updated = true
			break
		}
	}

	if !updated {
		fmt.Println("Task not found with ID:", id)
		return
	}

	saveTasks()
	fmt.Println("Task updated successfully.")
}

func PrintAllTasks() {
	loadTasks()

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("📋 Your Tasks:")
	for _, task := range tasks {
		fmt.Printf("🔹 ID: %s\n", task.ID)
		fmt.Printf("   Name: %s\n", task.Name)
		fmt.Printf("   Description: %s\n", task.Description)
		fmt.Printf("   Status: %s\n", task.Status)
		fmt.Printf("   Created At: %s\n", task.CreatedAt)
		fmt.Printf("   Updated At: %s\n", task.UpdatedAt)
		fmt.Println(strings.Repeat("-", 40))
	}
}

func PrintTasksByStatus(filterStatus string) {
	loadTasks()

	matched := false
	for _, task := range tasks {
		if string(task.Status) == filterStatus {
			fmt.Printf("🔹 ID: %s\n", task.ID)
			fmt.Printf("   Name: %s\n", task.Name)
			fmt.Printf("   Description: %s\n", task.Description)
			fmt.Printf("   Status: %s\n", task.Status)
			fmt.Printf("   Created At: %s\n", task.CreatedAt)
			fmt.Printf("   Updated At: %s\n", task.UpdatedAt)
			fmt.Println(strings.Repeat("-", 40))
			matched = true
		}
	}

	if !matched {
		fmt.Printf("No tasks found with status '%s'.\n", filterStatus)
	}
}
