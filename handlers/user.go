package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/pariyafesahat/go-rest-api/database"
	"github.com/pariyafesahat/go-rest-api/models"
)

var users = []models.User{
	{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   25,
	},
	{
		ID:    2,
		Name:  "Bob",
		Email: "bob@example.com",
		Age:   30,
	},
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	idString := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				fmt.Println("JSON encoding error:", err)
			}
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func GetUsers(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	users, err := database.GetUsers(conn)
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	createdUser, err := database.CreateUser(conn, user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(createdUser)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
	}
}
