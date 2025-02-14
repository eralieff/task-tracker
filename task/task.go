package task

import (
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(id int, description string) Task {
	return Task{
		Id:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func UpdateTask(tasks []Task, id int, description string) bool {
	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

func DeleteTask(tasks []Task, id int) ([]Task, bool) {
	for i := range tasks {
		if tasks[i].Id == id {
			return append(tasks[:i], tasks[i+1:]...), true
		}
	}
	return tasks, false
}
