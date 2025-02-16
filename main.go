package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"task-tracker/task"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run main.go <command> [args...]")
	}

	command := os.Args[1]
	tasks, err := task.ReadTasks()
	if err != nil {
		log.Fatal("error reading tasks: ", err)
	}

	switch command {
	case "add":
		if len(os.Args) != 3 {
			log.Fatal("usage: add <description>")
		}

		id := 1
		if len(tasks) > 0 {
			id = tasks[len(tasks)-1].Id + 1
		}

		newTask := task.NewTask(id, os.Args[2])
		tasks = append(tasks, newTask)

		if err := task.WriteTasks(tasks); err != nil {
			log.Fatal("error writing tasks: ", err)
		}
		fmt.Println("Task added.")

	case "update":
		if len(os.Args) != 4 {
			log.Fatal("usage: update <id> <description>")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("invalid task ID")
		}

		if updated := task.UpdateTask(tasks, id, os.Args[3]); !updated {
			log.Printf("task %d not found\n", id)
			return
		}

		if err = task.WriteTasks(tasks); err != nil {
			log.Fatal("error writing tasks: ", err)
		}
		fmt.Println("Task updated.")

	case "delete":
		if len(os.Args) != 3 {
			log.Fatal("usage: delete <id>")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("invalid task ID")
		}

		tasks, deleted := task.DeleteTask(tasks, id)
		if !deleted {
			log.Printf("task %d not found\n", id)
			return
		}

		if err = task.WriteTasks(tasks); err != nil {
			log.Fatal("error writing tasks: ", err)
		}
		fmt.Println("Task deleted.")

	case "mark-in-progress":
		{
			if len(os.Args) != 3 {
				log.Fatal("usage: mark-in-progress <id>")
			}

			err = mark(tasks, "in-progress")
			if err != nil {
				log.Fatal("error marking task: ", err)
			}

			fmt.Println("Task marked in-progress.")
		}
	case "mark-done":
		{
			if len(os.Args) != 3 {
				log.Fatal("usage: mark-done <id>")
			}

			err = mark(tasks, "done")
			if err != nil {
				log.Fatal("error marking task: ", err)
			}

			fmt.Println("Task marked done.")
		}
	case "list":
		{
			if len(os.Args) > 2 {
				switch os.Args[2] {
				case "done":
					{
						list(tasks, "done")
					}
				case "to-do":
					{
						list(tasks, "to-do")
					}
				case "in-progress":
					{
						list(tasks, "in-progress")
					}
				default:
					{
						log.Fatal("usage: list <done> <to-do> <in-progress>")
					}
				}
			} else {
				list(tasks, "")
			}
		}
	default:
		{
			fmt.Println("unknown command")
		}
	}
}

func mark(tasks []task.Task, status string) error {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return err
	}

	if marked := task.MarkTask(tasks, id, status); !marked {
		return fmt.Errorf("task %d not found", id)
	}

	if err = task.WriteTasks(tasks); err != nil {
		return err
	}

	return nil
}

func list(tasks []task.Task, status string) {
	for i := range tasks {
		if tasks[i].Status == status || status == "" {
			fmt.Println("ID:", tasks[i].Id)
			fmt.Println("Description:", tasks[i].Description)
			fmt.Println("Status:", tasks[i].Status)
			fmt.Println("Created:", tasks[i].CreatedAt)
			fmt.Println("Updated:", tasks[i].UpdatedAt)
			fmt.Println("")
		}
	}
}
