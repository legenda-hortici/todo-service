package handlers

import (
	"errors"
	"log"
	"todo-service/internal/services"
	"todo-service/internal/storage/domain"
	"todo-service/pkg/dto"

	"github.com/gofiber/fiber/v2"
)

type TasksHandler struct {
	service services.ITasksService
}

func NewTasksHandler(service services.ITasksService) *TasksHandler {
	return &TasksHandler{
		service: service,
	}
}

func (h *TasksHandler) CreateTask(c *fiber.Ctx) error {

	const (
		op = "handlers.tasks_handler.CreateTask"
	)

	var taskDTO dto.TaskDTO

	if err := c.BodyParser(&taskDTO); err != nil {
		log.Printf("%s: error parsing request body: %v", op, err)
		return c.JSON(fiber.Map{
			"code":  fiber.StatusBadRequest,
			"error": errors.New("error parsing request body"),
		})
	}

	task := domain.Task{
		Title:       taskDTO.Title,
		Description: taskDTO.Description,
	}

	taskID, err := h.service.CreateTask(task)
	if err != nil {
		log.Printf("%s: error creating task: %v", op, err)
		return c.JSON(fiber.Map{
			"code":  fiber.StatusInternalServerError,
			"error": errors.New("error creating task"),
		})
	}

	return c.JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "task created successfully",
		"id":      taskID,
	})
}

func (s *TasksHandler) GetTasks(c *fiber.Ctx) error {

	const (
		op = "handlers.tasks_handler.GetTasks"
	)

	tasks, err := s.service.GetTasks()
	if err != nil {
		log.Printf("%s: error getting tasks: %v", op, err)
		return c.JSON(fiber.Map{
			"code":  fiber.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code":      fiber.StatusOK,
		"all tasks": tasks,
	})
}

func (s *TasksHandler) UpdateTask(c *fiber.Ctx) error {

	const (
		op = "handlers.tasks_handler.UpdateTask"
	)

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("%s: error getting task id: %v", op, err)
		return c.JSON(fiber.Map{
			"code":  fiber.StatusBadRequest,
			"error": errors.New("error getting task id"),
		})
	}

	var taskDTO dto.TaskDTO
	if err := c.BodyParser(&taskDTO); err != nil {
		log.Printf("%s: error parsing request body: %v", op, err)
		return c.JSON(fiber.Map{
			"code":  fiber.StatusBadRequest,
			"error": errors.New("error parsing request body"),
		})
	}

	task := domain.Task{
		Title:       taskDTO.Title,
		Description: taskDTO.Description,
		Status:      taskDTO.Status,
	}

	err = s.service.UpdateTask(id, task)
	if err != nil {
		log.Printf("%s: error updating task: %v", op, err)
		return c.JSON(fiber.Map{
			"code":  fiber.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "task updated successfully",
	})
}

func (s *TasksHandler) DeleteTask(c *fiber.Ctx) error {

	const (
		op = "handlers.tasks_handler.DeleteTask"
	)

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("%s: error getting task id: %v", op, err)
		return c.JSON(fiber.Map{
			"code":  fiber.StatusBadRequest,
			"error": err.Error(),
		})
	}

	err = s.service.DeleteTask(id)
	if err != nil {
		log.Printf("%s: error deleting task: %v", op, err)
		return c.JSON(fiber.Map{
			"code":  fiber.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "task deleted successfully",
	})
}
