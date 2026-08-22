package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pariyafesahat/go-rest-api/database"
	"github.com/pariyafesahat/go-rest-api/handlers"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		fmt.Println("Databae Connection Error", err)
		return
	}
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
		if r.Method == http.MethodGet {
			handlers.GetUser(w, r, db)
			return
		}
		http.Error(w, "Method not alloed", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
