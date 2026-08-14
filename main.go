package main

import (
	"fmt"
	"net/http"

	"github.com/pariyafesahat/go-rest-api/handlers"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetUsers(w, r)
		case http.MethodPost:
			handlers.CreateUsers(w, r)
		default:
			http.Error(w, "Method not alloed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetUser(w, r)
			return
		}
		http.Error(w, "Method not alloed", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
