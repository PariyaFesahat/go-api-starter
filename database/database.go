package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// connString := "postgres://postgres:admin@localhost:5432/go_rest_api"
func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:admin@localhost:5432/go_rest_api",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return conn, nil
}
