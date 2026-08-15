package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {

	connString := "postgres://postgres:admin@localhost:5432/go_rest_api"
	conn, err := pgx.Connect(
		context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return conn, nil
}
