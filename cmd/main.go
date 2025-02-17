package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"task-tracker/internal/repository"
	"task-tracker/internal/usecase"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run cmd/main.go <command> [args...]")
	}

	repo := &repository.FileTaskRepository{}
	service := &usecase.TaskService{Repo: repo}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) != 3 {
			log.Fatal("usage: add <description>")
		}

		err := service.AddTask(os.Args[2])
		handleError(err, "Task added.")

	case "update":
		if len(os.Args) != 4 {
			log.Fatal("usage: update <id> <description>")
		}

		id, err := strconv.Atoi(os.Args[2])
		handleError(err, "")

		err = service.UpdateTask(id, os.Args[3])
		handleError(err, "Task updated.")

	case "delete":
		if len(os.Args) != 3 {
			log.Fatal("usage: delete <id>")
		}

		id, err := strconv.Atoi(os.Args[2])
		handleError(err, "")

		err = service.DeleteTask(id)
		handleError(err, "Task deleted.")

	case "mark-done", "mark-in-progress":
		if len(os.Args) != 3 {
			log.Fatal("usage: mark-(in-progress or done) <id>")
		}

		id, err := strconv.Atoi(os.Args[2])
		handleError(err, "")

		status := strings.TrimPrefix(os.Args[1], "mark-")

		err = service.MarkTask(id, status)
		handleError(err, "Task status updated.")

	case "list":
		if len(os.Args) > 3 {
			log.Fatal("usage: list (optional <done | todo | in-progress>)")
		}

		status := ""
		if len(os.Args) == 3 {
			validStatuses := map[string]bool{"done": true, "todo": true, "in-progress": true}
			if !validStatuses[os.Args[2]] {
				log.Fatal("Invalid status. usage: list (optional <done | todo | in-progress>)")
			}
			status = os.Args[2]
		}

		tasks, err := service.ListTasks(status)
		handleError(err, "")

		for _, t := range tasks {
			fmt.Printf("ID: %d, Desc: %s, Status: %s\n", t.Id, t.Description, t.Status)
		}

	default:
		log.Fatal("unknown command")
	}
}

func handleError(err error, successMsg string) {
	if err != nil {
		log.Fatal(err)
	} else if successMsg != "" {
		fmt.Println(successMsg)
	}
}
