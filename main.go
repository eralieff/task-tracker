package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func main() {
	// To run the program you need to write the following in the command line: “go run main.go <args...>”
	switch os.Args[1] {
	case "add":
		{
			fmt.Println("add")

			if len(os.Args) != 3 {
				log.Fatal("usage: add <description>")
			}

			data, err := readFileIfExists("tasks.json")
			if err != nil {
				log.Fatal("error reading file: ", err)
			}

			var tasks []Task
			maxId := 0
			if len(data) != 0 {
				err = json.Unmarshal(data, &tasks)
				if err != nil {
					log.Fatal("error unmarshalling json: ", err)
				}

				maxId = tasks[len(tasks)-1].Id
			}

			task := Task{
				Id:          maxId + 1,
				Description: os.Args[2],
				Status:      "todo",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			tasks = append(tasks, task)

			err = writeTasksToFile(tasks, "tasks.json")
			if err != nil {
				log.Fatal("error writing tasks to file: ", err)
			}
		}
	case "update":
		{
			fmt.Println("update")
			if len(os.Args) != 4 {
				log.Fatal("usage: update <id> <description>")
			}

			data, err := readFileIfExists("tasks.json")
			if err != nil {
				log.Fatal("error reading file: ", err)
			}

			var tasks []Task
			if len(data) != 0 {
				err = json.Unmarshal(data, &tasks)
				if err != nil {
					log.Fatal("error unmarshalling json: ", err)
				}
			} else {
				log.Println("file is empty, you need to add tasks")
				return
			}

			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("error converting id to int: ", err)
			}

			taskExists := false
			for i := range tasks {
				if tasks[i].Id == id {
					tasks[i].Description = os.Args[3]
					taskExists = true
					break
				}
			}

			if !taskExists {
				log.Printf("task %d not found", id)
				return
			}

			err = writeTasksToFile(tasks, "tasks.json")
			if err != nil {
				log.Fatal("error writing tasks to file: ", err)
			}
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

func readFileIfExists(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("file does not exist")
			return nil, nil
		}
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal("error closing file: ", err)
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func writeTasksToFile(tasks []Task, filename string) error {
	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal("error closing file: ", err)
		}
	}()

	_, err = io.WriteString(file, string(bytes))
	if err != nil {
		return err
	}

	return nil
}
