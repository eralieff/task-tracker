package main

import (
	"fmt"
	"os"
)

func main() {
	// To run the program you need to write the following in the command line: “go run main.go <args...>”
	switch os.Args[1] {
	case "add":
		{
			fmt.Println("add")
		}
	case "update":
		{
			fmt.Println("update")
		}
	case "delete":
		{
			fmt.Println("delete")
		}
	case "mark-in-progress":
		{
			fmt.Println("mark-in-progress")
		}
	case "mark-done":
		{
			fmt.Println("mark-done")
		}
	case "list":
		{
			if len(os.Args) > 2 {
				switch os.Args[2] {
				case "done":
					{
						fmt.Println("done")
					}
				case "to-do":
					{
						fmt.Println("to-do")
					}
				case "in-progress":
					{
						fmt.Println("in-progress")
					}
				default:
					{
						fmt.Println("unknown command")
					}
				}
			} else {
				fmt.Println("list")
			}
		}
	default:
		{
			fmt.Println("unknown command")
		}
	}
}
