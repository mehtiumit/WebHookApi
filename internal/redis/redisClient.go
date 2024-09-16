package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"webhook/internal/domain/entities"
	"webhook/pkg/log"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value entities.HookRedisModel) error
}

type redisClient struct {
	client *redis.Client
	logger *log.Logrus
}

func (r redisClient) Set(ctx context.Context, key string, value entities.HookRedisModel) error {
	jsonContent, _ := json.Marshal(value)
	err := r.client.Set(ctx, key, jsonContent, 0).Err()
	if err != nil {
		r.logger.Errorf("Error while setting data to redis. | Error: %s", err.Error())
		return err
	}
	return nil
}

func NewRedisClient(logger *log.Logrus) RedisClient {
	var rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	return &redisClient{client: rdb, logger: logger}
}
