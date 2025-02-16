package repository

import (
	"encoding/json"
	"os"
	"task-tracker/internal/entity"
)

const FileName = "tasks.json"

type TaskRepository interface {
	ReadTasks() ([]entity.Task, error)
	WriteTasks([]entity.Task) error
}

type FileTaskRepository struct{}

func (r *FileTaskRepository) ReadTasks() ([]entity.Task, error) {
	file, err := os.Open(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []entity.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *FileTaskRepository) WriteTasks(tasks []entity.Task) error {
	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(FileName, bytes, 0644)
}
