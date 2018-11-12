package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"github.com/sirkon/weather-cacher/internal/schema"
	"time"
)

// Redis конструктор кеша работающего поверх redis-а
func Redis(c *redis.Client) Cache {
	return &redisImpl{
		client: c,
		prefix: "redis-cache::",
	}
}

type redisImpl struct {
	client *redis.Client
	prefix string
}

func (cr *redisImpl) keyName(key string) string {
	return cr.prefix + key
}

// Get ...
func (cr *redisImpl) Get(ctx context.Context, forecastID string) (*schema.Forecast, error) {
	res := cr.client.WithContext(ctx).Get(cr.keyName(forecastID))
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("failed to query for a given key: %s", err)
	}

	result := &schema.Forecast{}
	data, err := res.Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to get a value bound to given key: %s", err)
	}

	if err := proto.Unmarshal(data, result); err != nil {
		return nil, fmt.Errorf("got corrupted data from source: %s", err)
	}

	return result, nil
}

// Set ...
func (cr *redisImpl) Set(ctx context.Context, forecastID string, forecast *schema.Forecast) error {
	data, err := proto.Marshal(forecast)
	if err != nil {
		return fmt.Errorf("failed to serialize forecast structure: %s", err)
	}

	if err := cr.client.WithContext(ctx).Set(cr.keyName(forecastID), data, time.Minute*15).Err(); err != nil {
		return fmt.Errorf("failed to store forecast value: %s", err)
	}

	return nil
}
