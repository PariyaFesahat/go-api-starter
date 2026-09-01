package main

// @title Go REST API
// @version 1.0
// @description A REST API built with Go, PostgreSQL, Redis, and Docker.
//
// @host localhost:8080

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pariyafesahat/go-rest-api/database"
	"github.com/pariyafesahat/go-rest-api/handlers"
	"github.com/pariyafesahat/go-rest-api/redis"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Database Connection Error", err)
		return
	}

	redis.Connect()

	defer db.Close(context.Background())

	fmt.Println("Connected to PostgreSQL successfully")

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetUsers(w, r, db)
		case http.MethodPost:
			handlers.CreateUser(w, r, db)
		default:
			http.Error(w, "Method not alloed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetUser(w, r, db)
		case http.MethodPut:
			handlers.UpdateUser(w, r, db)
		case http.MethodDelete:
			handlers.DeleteUser(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
