package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID          int
	Title       string
	IsCompleted bool
}

type TaskManager struct {
	tasks map[int]Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[int]Task),
	}
}

func (tm *TaskManager) AddTask(title string) {
	id := len(tm.tasks) + 1
	task := Task{
		ID:          id,
		Title:       title,
		IsCompleted: false,
	}
	tm.tasks[id] = task
}

func (tm *TaskManager) RemoveTask(id int) bool {
	_, ok := tm.tasks[id]

	if !ok {
		fmt.Print("Task not found\n")
		return ok
	}
	delete(tm.tasks, id)
	return ok
}

func (tm *TaskManager) CompleteTask(id int) bool {

	task, ok := tm.tasks[id]
	if !ok {
		fmt.Print("Task not found\n")
		return ok
	}
	task.IsCompleted = true
	tm.tasks[id] = task
	return ok
}

func (tm *TaskManager) PrintTasks() {
	fmt.Print("\nList of tasks:\n")

	for _, task := range tm.tasks {
		fmt.Printf("ID: %d, Title: %s, Completed: %t\n", task.ID, task.Title, task.IsCompleted)
	}
}

func TaskManagerHandler(taskManager *TaskManager) {
	var currentNumberAction int

	fmt.Print("\nActions \n[1] Add new task\n[2] Delete task\n[3] View list of tasks\n[4] Set task as complited\nPlease, enter a number: ")
	_, err := fmt.Scanln(&currentNumberAction)
	if err != nil {
		fmt.Print("Invalid action\n")
		fmt.Scanln()
		return
	}

	if len(taskManager.tasks) == 0 && currentNumberAction != 1 {
		fmt.Println("\nNo tasks to display ")
		return
	}

	switch currentNumberAction {
	case 1:
		var title string

		fmt.Print("Enter task title: ")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter task title: ")
		title, err = reader.ReadString('\n')
		title = strings.TrimSpace(title)
		if title == "" {
			fmt.Print("\nTitle cannot be empty\n")
			return
		}

		taskManager.AddTask(title)
		fmt.Println("\nTask added successfully.")
	case 2:
		var id int

		fmt.Print("Enter task ID to delete: ")
		_, err := fmt.Scanln(&id)
		if err != nil {
			fmt.Print("Invalid id\n")
		}
		if deleted := taskManager.RemoveTask(id); deleted {
			fmt.Println("\nTask deleted successfully.")
		}
	case 3:
		taskManager.PrintTasks()
	case 4:
		var id int

		fmt.Print("Enter task ID to mark as completed: ")
		_, err := fmt.Scanln(&id)
		if err != nil {
			fmt.Println("Invalid id")
		}

		if checked := taskManager.CompleteTask(id); checked {
			fmt.Println("\nTask completed successfully.")
		}
	default:
		return
	}

}

func main() {
	taskManager := NewTaskManager()

	for {
		TaskManagerHandler(taskManager)
	}

}
