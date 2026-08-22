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

func GetUser(conn *pgx.Conn, id int) (models.User, error) {
	var user models.User

	err := conn.QueryRow(
		context.Background(),
		`
		SELECT id, name, email, age
		FROM users
		WHERE id = $1
		`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Age,
	)

	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
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

	users := make([]models.User, 0)

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

func CreateUser(conn *pgx.Conn, user models.User) (models.User, error) {
	err := conn.QueryRow(
		context.Background(),
		`
		INSERT INTO users (name, email, age)
		VALUES ($1, $2, $3)
		RETURNING id
		`,
		user.Name,
		user.Email,
		user.Age,
	).Scan(&user.ID)

	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
