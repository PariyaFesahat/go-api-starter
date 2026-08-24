package redis

import (
	"context"
	"log"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

var Client *goredis.Client

func Connect() {
	Client = goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379",
	})

	err := Client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Connected to Redis")
}

// Set stores a value in Redis.
func Set(key string, value string, expiration time.Duration) error {
	return Client.Set(context.Background(), key, value, expiration).Err()
}

// Get retrieves a value from Redis.
func Get(key string) (string, error) {
	return Client.Get(context.Background(), key).Result()
}

// Delete removes a value from Redis.
func Delete(key string) error {
	return Client.Del(context.Background(), key).Err()
}
