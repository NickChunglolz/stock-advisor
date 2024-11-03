package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
)

const (
	POSTGRES_HOST = "POSTGRES_HOST"
	POSTGRES_PORT = "POSTGRES_PORT"
	REDIS_HOST    = "REDIS_HOST"
	REDIS_PORT    = "REDIS_PORT"
)

var (
	postgresClient *pg.DB
	postgresOnce   sync.Once
	redisClient    *redis.Client
	redisOnce      sync.Once
)

type DatabaseFactoryInterface interface {
	CreatePostgres() (*pg.DB, error)
	CreateRedis() (*redis.Client, error)
	GetPostgres() *pg.DB
	GetRedis() *redis.Client
	CloseDatabaseConnections()
}

type DatabaseFactory struct{}

func NewDatabaseFactory() *DatabaseFactory {
	return &DatabaseFactory{}
}

func (df *DatabaseFactory) CreatePostgres() (*pg.DB, error) {
	var err string

	postgresOnce.Do(func() {
		postgresClient = pg.Connect(&pg.Options{
			Addr: fmt.Sprintf("%s:%s", os.Getenv(POSTGRES_HOST), os.Getenv(POSTGRES_PORT)),
			User: "postgres",
			Password: "postgres",
		})

		// err = postgresClient.Ping(context.Background()).Error()
		// if len(err) > 0 {
		// 	log.Println("Error connecting to Postgres:", err)
		// }
	})

	if len(err) > 0 {
		return nil, errors.New(err)
	}

	fmt.Println("Postgres client created successfully:", postgresClient)
	return postgresClient, nil
}

func (df *DatabaseFactory) CreateRedis() (*redis.Client, error) {
	var err error

	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", os.Getenv(REDIS_HOST), os.Getenv(REDIS_PORT)),
		})
		_, err = redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Println("Error connecting to Redis:", err)
		}
	})

	if err != nil {
		return nil, err
	} else {
		log.Println("Redis client created successfully:", redisClient)
		return redisClient, nil
	}
}

func (df *DatabaseFactory) GetPostgres() *pg.DB {
	return postgresClient
}

func (df *DatabaseFactory) GetRedis() *redis.Client {
	return redisClient
}

func (df *DatabaseFactory) CloseDatabaseConnections() {
	if redisClient != nil {
		redisClient.Close()
	}
}
