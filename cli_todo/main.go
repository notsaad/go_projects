package main

import (
	"flag"
	"fmt"
	"os"
)

type Task struct {
	ID     string `json:"id"`
	text   string `json:"text"`
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
	case "list":
		fmt.Println("Listing tasks")
	}

}
