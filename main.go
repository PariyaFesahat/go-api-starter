package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pariyafesahat/go-rest-api/database"
	"github.com/pariyafesahat/go-rest-api/handlers"
)

func main() {
	conn, err := database.Connect()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}
	defer conn.Close(context.Background())

	fmt.Println("Connected to PostgreSQL!")

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.Handle("/users", handlers.GetUsers(conn))

	fmt.Println("Server is running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
