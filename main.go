package main

import (
	"flag"
	"fmt"
	data "todoList/datbase"
	opration "todoList/oprations"
)

func main() {
	// cmd.Execute()
	data.ConnectDB()

	command := flag.String("command", "", "Command to execute: add, edit, update, list")
	task := flag.String("task", "", "The task to add, edit, or update")
	qty := flag.Int("qty", 0, "The quantity of the item")
	listName := flag.String("list", "Default List", "The name of the list")
	id := flag.Int("id", 0, "listID")
	flag.Parse()

	switch *command {

	case "add":
		if *task == "" {
			fmt.Println("Please provide a task using the -task flag.")
			return
		}

		item := opration.Items{
			Name: *task,
			Qty:  *qty,
		}

		opration.Add(*listName, []opration.Items{item})

	case "update":

		if *task == "" {
			fmt.Println("Please provide a task using the -task flag.")
			return
		}

		item := opration.Items{
			Name: *task,
			Qty:  *qty,
		}

		opration.Update(27, *listName, []opration.Items{item})

	case "delete":

		opration.Delete(*id)
	}

}
