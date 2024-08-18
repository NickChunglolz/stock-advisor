package utils

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

type DatabaseFactoryInterface interface {
	CreatePostgres() error
	CreateRedis() (*redis.Client, error)
	GetPostgres()
	GetRedis() *redis.Client
	CloseDatabaseConnections()
}

type DatabaseFactory struct{}

func (df *DatabaseFactory) CreatePostgres() error {
	// Implement PostgreSQL client initialization here
	return nil
}

func (df *DatabaseFactory) CreateRedis() *redis.Client {
	var err error

	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
		_, err = redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Println("Error connecting to Redis:", err)
		}
	})

	fmt.Println("Redis client created successfully:", redisClient)

	return redisClient
}

func (df *DatabaseFactory) GetRedis() *redis.Client {
	return redisClient
}

func (df *DatabaseFactory) CloseDatabaseConnections() {
	if redisClient != nil {
		redisClient.Close()
	}
}
