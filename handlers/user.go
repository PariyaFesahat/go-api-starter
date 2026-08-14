package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pariyafesahat/go-rest-api/models"
)

var users = []models.User{
	{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   22,
	},
	{
		Name:  "Bob",
		Email: "bob@example.com",
		Age:   30,
	},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
	}
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
	}

}
