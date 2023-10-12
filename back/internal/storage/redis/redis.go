package redis

import "github.com/redis/go-redis/v9"

type Storage struct {
	client *redis.Client
}

func New() *Storage {
	client := redis.NewClient(&redis.Options{
		Addr:     "network_with_redis:6379",
		Password: "",
		DB:       0,
	})

	return &Storage{
		client: client,
	}
}
