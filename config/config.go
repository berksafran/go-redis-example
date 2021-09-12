package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type config struct {
	Client      *redis.Client
	Ctx         context.Context
	ChannelName string
}

// Config includes context and redis client
var Config config
