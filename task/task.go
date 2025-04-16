package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

const dataFile = "tasks.json"

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(dataFile, data, 0644)
	return err
}

func AddTask(title, description string) error {
	tasks, err := loadTasks()
	if err != nil {
		tasks = []Task{}
	}
	newTask := Task{
		ID:          len(tasks) + 1,
		Title:       title,
		Description: description,
		Done:        false,
	}
	tasks = append(tasks, newTask)
	return saveTasks(tasks)
}
func UpdateTask(id int, title, description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = title
			tasks[i].Description = description
			return saveTasks(tasks)
		}
	}
	return nil
}

func ListTasks() error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Sprintf("tasks not found")
		return nil
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Done: %t\n", task.ID, task.Title, task.Description, task.Done)
	}
	return nil
}

func DoneTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Sprintf("Task with ID %d marked as done\n", id)
			return saveTasks(tasks)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
func DeleteTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return saveTasks(tasks)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
