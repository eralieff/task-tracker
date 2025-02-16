package usecase

import (
	"errors"
	"task-tracker/internal/entity"
	"task-tracker/internal/repository"
	"time"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func (s *TaskService) AddTask(description string) error {
	tasks, err := s.Repo.ReadTasks()
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].Id + 1
	}

	newTask := entity.NewTask(id, description)
	tasks = append(tasks, newTask)

	return s.Repo.WriteTasks(tasks)
}

func (s *TaskService) UpdateTask(id int, description string) error {
	tasks, err := s.Repo.ReadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return s.Repo.WriteTasks(tasks)
		}
	}

	return errors.New("task not found")
}

func (s *TaskService) DeleteTask(id int) error {
	tasks, err := s.Repo.ReadTasks()
	if err != nil {
		return err
	}

	newTasks := []entity.Task{}
	found := false
	for _, t := range tasks {
		if t.Id != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("task not found")
	}

	return s.Repo.WriteTasks(newTasks)
}

func (s *TaskService) MarkTask(id int, status string) error {
	tasks, err := s.Repo.ReadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return s.Repo.WriteTasks(tasks)
		}
	}

	return errors.New("task not found")
}

func (s *TaskService) ListTasks(status string) ([]entity.Task, error) {
	tasks, err := s.Repo.ReadTasks()
	if err != nil {
		return nil, err
	}

	if status == "" {
		return tasks, nil
	}

	filtered := []entity.Task{}
	for _, task := range tasks {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}
