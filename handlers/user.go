package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pariyafesahat/go-rest-api/database"
	"github.com/pariyafesahat/go-rest-api/models"
	"github.com/pariyafesahat/go-rest-api/redis"
)

func invalidateUsersCache() {
	err := redis.Delete("users:all")
	if err != nil {
		fmt.Println("Redis users cache delete error:", err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	idString := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cacheKey := fmt.Sprintf("user:%d", id)

	// Try Redis first.
	cachedUser, err := redis.Get(cacheKey)
	if err == nil {
		var user models.User

		err = json.Unmarshal([]byte(cachedUser), &user)
		if err == nil {
			fmt.Println("Cache hit")

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(user)
			if err != nil {
				fmt.Println("JSON encoding error:", err)
			}
			return
		}
	}

	fmt.Println("Cache miss")

	// Cache miss: get the user from PostgreSQL.
	user, err := database.GetUser(conn, id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Store the user in Redis for 5 minutes.
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
	} else {
		err = redis.Set(cacheKey, string(userJSON), 5*time.Minute)
		if err != nil {
			fmt.Println("Redis cache error:", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	cacheKey := "users:all"

	// Try Redis first.
	cachedUsers, err := redis.Get(cacheKey)
	if err == nil {
		var users []models.User

		err = json.Unmarshal([]byte(cachedUsers), &users)
		if err == nil {
			fmt.Println("Users cache hit")

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(users)
			if err != nil {
				fmt.Println("JSON encoding error:", err)
			}
			return
		}
	}

	fmt.Println("Users cache miss")

	// Cache miss: get users from PostgreSQL.
	users, err := database.GetUsers(conn)
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	// Store users in Redis for 5 minutes.
	usersJSON, err := json.Marshal(users)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
	} else {
		err = redis.Set(cacheKey, string(usersJSON), 5*time.Minute)
		if err != nil {
			fmt.Println("Redis cache error:", err)
		}
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

	// A new user changes the full users list.
	invalidateUsersCache()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(createdUser)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	idString := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	updatedUser, err := database.UpdateUser(conn, id, user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Remove the individual user's cached data.
	cacheKey := fmt.Sprintf("user:%d", id)

	err = redis.Delete(cacheKey)
	if err != nil {
		fmt.Println("Redis cache delete error:", err)
	}

	// The updated user also affects the full users list.
	invalidateUsersCache()

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(updatedUser)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	idString := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteUser(conn, id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Remove the individual user's cached data.
	cacheKey := fmt.Sprintf("user:%d", id)

	err = redis.Delete(cacheKey)
	if err != nil {
		fmt.Println("Redis cache delete error:", err)
	}

	// A deleted user also changes the full users list.
	invalidateUsersCache()

	w.WriteHeader(http.StatusNoContent)
}
