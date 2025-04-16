package main

import (
	"fmt"
	"go-todo/task"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Gunakan todo [add][list][done][delete]")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		//func add()
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
		fmt.Printf("list todo")
	case "done":
		//func done()
		fmt.Printf("done todo")
	case "delete":
		//func delete()
		fmt.Printf("delete todo")
	default:
		fmt.Printf("Gunakan todo [add][list][done][delete]")

	}
}
