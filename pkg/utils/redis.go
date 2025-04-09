package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisAuthStore struct {
	clientAuth *redis.Client
}

func NewRedisAuthStore(redisAddr string) *RedisAuthStore {
	var client *redis.Client

	for retries := 0; retries < 5; retries++ {
		client = redis.NewClient(&redis.Options{
			Addr: redisAddr,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		status := client.Ping(ctx)
		if status.Err() == nil {
			fmt.Println("Connected to Redis successfully")
			store := &RedisAuthStore{clientAuth: client}
			go store.reconnect(redisAddr)
			return store
		}

		fmt.Printf("Failed to connect to Redis: %v. Retrying in 5 seconds...", status.Err())
		time.Sleep(5 * time.Second)
	}

	fmt.Printf("Failed to connect to Redis after 5 attempts")
	return nil
}

func (store *RedisAuthStore) reconnect(redisAddr string) {
	for {
		time.Sleep(30 * time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		status := store.clientAuth.Ping(ctx)
		cancel()
		if status.Err() != nil {
			fmt.Printf("Lost connection to Redis: %v. Attempting to reconnect...", status.Err())
			for retries := 0; retries < 5; retries++ {
				client := redis.NewClient(&redis.Options{
					Addr: redisAddr,
				})

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				status := client.Ping(ctx)
				cancel()
				if status.Err() == nil {
					fmt.Println("Reconnected to Redis successfully")
					store.clientAuth = client
					break
				}

				fmt.Printf("Failed to reconnect to Redis: %v. Retrying in 5 seconds...", status.Err())
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func (s *RedisAuthStore) Get(key string) (string, error) {
	localCtx := context.Background()
	value, err := s.clientAuth.Get(localCtx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (s *RedisAuthStore) Set(key string, value interface{}) error {
	localCtx := context.Background()
	err := s.clientAuth.Set(localCtx, key, value, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}
