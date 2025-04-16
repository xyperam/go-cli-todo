package task

import (
	"encoding/json"
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
