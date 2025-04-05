package routes

import (
	"todo-service/internal/handlers"
	"todo-service/internal/services"
	"todo-service/internal/storage/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
)

func InitRoutes(port string, conn *pgx.Conn) *fiber.App {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	repo := repository.NewTasksRepository(conn)
	services := services.NewTasksService(repo)
	handlers := handlers.NewTasksHandler(services)

	api := app.Group("/api")
	{
		api.Post("/tasks", handlers.CreateTask)
		api.Get("/tasks", handlers.GetTasks)
		api.Put("/tasks/:id", handlers.UpdateTask)
		api.Delete("/tasks/:id", handlers.DeleteTask)
	}
		

	return app
}
