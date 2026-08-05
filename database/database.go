package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	dsn := "postgres://postgres:admin@localhost:5432/go_rest_api"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to PostgreSQL: %w", err)
	}

	return conn, nil
}
