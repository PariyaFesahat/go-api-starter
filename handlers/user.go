package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/pariyafesahat/go-rest-api/models"
)

func GetUsers(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := conn.Query(context.Background(), "SELECT id, name, email FROM users")
		if err != nil {
			http.Error(w, "Failed to get users", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []models.User

		for rows.Next() {
			var user models.User

			err := rows.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
			)
			if err != nil {
				http.Error(w, "Failed to read user", http.StatusInternalServerError)
				return
			}

			users = append(users, user)
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(users)
	}
}
