package services

import (
	"fmt"
	"todo-service/internal/storage/domain"
	"todo-service/internal/storage/repository"
)

type ITasksService interface {
	CreateTask(domain.Task) (int, error)
	GetTasks() ([]domain.Task, error)
	UpdateTask(int, domain.Task) error
	DeleteTask(int) error
}

type TasksService struct {
	repo repository.ITasksRepository
}

func NewTasksService(repo repository.ITasksRepository) *TasksService {
	return &TasksService{
		repo: repo,
	}
}

func (s *TasksService) CreateTask(task domain.Task) (int, error) {

	const (
		op = "services.tasks_service.CreateTask"
	)

	taskID, err := s.repo.CreateTask(task)
	if err != nil {
		return -1, fmt.Errorf("error creating task: %w: %v", err, op)
	}

	return taskID, nil
}

func (s *TasksService) GetTasks() ([]domain.Task, error) {

	const (
		op = "services.tasks_service.GetTasks"
	)

	tasks, err := s.repo.GetTasks()
	if err != nil {
		return []domain.Task{}, fmt.Errorf("error getting tasks: %w: %v", err, op)
	}

	return tasks, nil
}

func (s *TasksService) UpdateTask(id int, task domain.Task) error {

	const (
		op = "services.tasks_service.UpdateTask"
	)

	err := s.repo.UpdateTask(id, task)
	if err != nil {
		return fmt.Errorf("error updating task: %w: %v", err, op)
	}

	return nil
}

func (s *TasksService) DeleteTask(id int) error {

	const (
		op = "services.tasks_service.DeleteTask"
	)

	err := s.repo.DeleteTask(id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w: %v", err, op)
	}

	return nil
}
