package task

import (
	"encoding/json"
	"io"
	"os"
)

const FileName = "tasks.json"

func ReadTasks() ([]Task, error) {
	file, err := os.Open(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if len(data) > 0 {
		if err = json.Unmarshal(data, &tasks); err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

func WriteTasks(tasks []Task) error {
	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(FileName, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
