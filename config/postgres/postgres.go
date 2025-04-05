package postgres

import (
	"context"
	"fmt"
	"log"
	"todo-service/internal/config"

	"github.com/jackc/pgx/v5"
)

func InitDB(cfg *config.Config) (*pgx.Conn, error) {
	const (
		op = "storage.postgres"
	)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	log.Printf("%s: connected to database", op)

	return conn, nil
}
