package redis

import (
	"bluebell/settings"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

func Init(cfg settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	ctx := context.Background()
	_, err = client.Ping(ctx).Result()

	if err != nil {
		return
	}

	return
}

func Close() {
	_ = client.Close()
}
