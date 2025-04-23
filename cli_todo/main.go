package main

import (
	"flag"
	"fmt"
	"os"
)

type Task struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Status bool   `json:"status"`
}

func main() {
	// first parameter - name of the flag (what users type on the cli)
	// second parameter - default value that is used if the user doesnt specify a value
	// third parameter - usage string, appears when someone uses the help flag
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'list' subcommand")
		os.Exit(1) // general error code
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Println("Adding a task", *addCmd)
		exTask := Task{
			ID:     "1",
			Text:   "Learn Go",
			Status: false,
		}
		fmt.Println("Task Added:", exTask)
	case "list":
		fmt.Println("Listing tasks")
	default:
		fmt.Println("Unrecognized command")
		os.Exit(1) 
	}

	// 0644 is the permission of the file
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// calls f.Close() when the main function returns
	defer f.Close()

	f.WriteString("hello world\n")

}
