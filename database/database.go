package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pariyafesahat/go-rest-api/models"
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

func GetUsers(conn *pgx.Conn) ([]models.User, error) {
	rows, err := conn.Query(context.Background(), `
		SELECT id, name, email, age
		FROM users
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Age,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while reading users: %w", err)
	}

	return users, nil
}
