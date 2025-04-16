package main

import (
	"fmt"
	"go-todo/task"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Gunakan todo [add][list][done][delete]")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		//func add()
		if len(os.Args) < 4 {
			fmt.Printf("Gunakan todo add [title] [description]")
			return
		}
		fmt.Printf("add todo")
		title := os.Args[2]
		description := os.Args[3]
		err := task.AddTask(title, description)
		if err != nil {
			fmt.Printf("Error adding task: %v", err)
			return
		}
		fmt.Printf("Task added| title: %s  Description: %s", title, description)

	case "list":
		//func list()
		err := task.ListTasks()
		if err != nil {
			fmt.Printf("Error listing tasks: %v", err)
		}
		fmt.Printf("list todo")
	case "done":
		//func done()
		if len(os.Args) < 3 {
			fmt.Printf("Gunakan todo done [id]")
			return
		}
		idstr := os.Args[2]
		id, err := strconv.Atoi(idstr)
		if err != nil {
			fmt.Printf("Error converting id to int: %v", err)
			return
		}
		err = task.DoneTask(id)
		if err != nil {
			fmt.Printf("Error marking task as done: %v", err)
			return
		}
		fmt.Printf("Task with ID %d marked as done\n", id)
	case "delete":
		//func delete()
		if len(os.Args) < 3 {
			fmt.Printf("Gunakan todo delete [id]")
			return
		}
		idstr := os.Args[2]
		id, err := strconv.Atoi(idstr)
		if err != nil {
			fmt.Printf("Error converting id to int: %v", err)
			return
		}
		err = task.DeleteTask(id)
		if err != nil {
			fmt.Printf("Error deleting task: %v", err)
			return
		}
		fmt.Printf("Task with ID %d deleted\n", id)
	default:
		fmt.Printf("Gunakan todo [add][list][done][delete]")
	}
}
