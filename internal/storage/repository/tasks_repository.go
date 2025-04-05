package repository

import (
	"context"
	"fmt"
	"todo-service/internal/storage/domain"

	"github.com/jackc/pgx/v5"
)

type ITasksRepository interface {
	CreateTask(domain.Task) (int, error)
	GetTasks() ([]domain.Task, error)
	DeleteTask(int) error
	UpdateTask(int, domain.Task) error
}

type TaskRepository struct {
	conn *pgx.Conn
}

func NewTasksRepository(conn *pgx.Conn) *TaskRepository {
	return &TaskRepository{
		conn: conn,
	}
}

func (r *TaskRepository) CreateTask(task domain.Task) (int, error) {

	const (
		op = "storage.repository.tasks_repository.CreateTask"
	)

	err := r.conn.QueryRow(context.Background(),
		"INSERT INTO tasks (title, description, status, created_at, updated_at) VALUES ($1, $2, default, now(), now()) RETURNING id",
		task.Title, task.Description).Scan(&task.ID)
	if err != nil {
		return -1, fmt.Errorf("error creating task: %w: %v", err, op)
	}
	return task.ID, nil
}

func (r *TaskRepository) GetTasks() ([]domain.Task, error) {

	const (
		op = "storage.repository.tasks_repository.GetTasks"
	)

	rows, err := r.conn.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("error getting tasks: %w: %v", err, op)
	}
	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning task: %w: %v", err, op)
		}
		tasks = append(tasks, task)
	}
	rows.Close()

	return tasks, nil
}

func (r *TaskRepository) UpdateTask(id int, task domain.Task) error {

	const (
		op = "storage.repository.tasks_repository.UpdateTask"
	)

	_, err := r.conn.Exec(context.Background(),
		"UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = now() WHERE id = $4",
		task.Title, task.Description, task.Status, id)
	if err != nil {
		return fmt.Errorf("error updating task: %w: %v", err, op)
	}
	fmt.Println(err)

	return nil
}

func (r *TaskRepository) DeleteTask(id int) error {

	const (
		op = "storage.repository.tasks_repository.DeleteTask"
	)

	_, err := r.conn.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w: %v", err, op)
	}

	return nil
}
