package main

import (
	"context"
	"go-redis-example/cmd"
	"go-redis-example/config"

	"github.com/go-redis/redis/v8"
)

func init() {
	// Create a context
	var ctx = context.Background()

	// Create a new client for Redis
	// Ensure that redis-server is running on 6379 port
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // or 127.0.0.1:6379
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	// Set context and Redis client on Config
	config.Config.Client = rdb
	config.Config.Ctx = ctx
}

func main() {
	cmd.Execute()
}
