package server

import (
	"context"
	"log"
	"todo-service/cmd/routes"
	"todo-service/config/postgres"
	"todo-service/internal/config"
)

func RunApp() {
	cfg := config.InitConfig()
	log.Println("config loaded successfully")

	conn, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
	log.Println("successfully connected to database")

	app := routes.InitRoutes(cfg.HTTPServer.Port, conn)

	if err := app.Listen(":" + cfg.HTTPServer.Port); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
