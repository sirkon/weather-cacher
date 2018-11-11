package weather

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// CacheRedis конструктор кеша работающего поверх redis-а
func CacheRedis(c *redis.Client) Cache {
	return &cacheRedis{
		client: c,
		prefix: "redis-cache::",
	}
}

type cacheRedis struct {
	client *redis.Client
	prefix string
}

func (cr *cacheRedis) keyName(key string) string {
	return cr.prefix + key
}

// Get ...
func (cr *cacheRedis) Get(ctx context.Context, forecastID string) (*Forecast, error) {
	res := cr.client.WithContext(ctx).Get(cr.keyName(forecastID))
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("failed to query for a given key: %s", err)
	}

	result := &Forecast{}
	data, err := res.Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to get a value bound to given key: %s", err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("got corrupted data from source: %s", err)
	}

	return result, nil
}

// Set ...
func (cr *cacheRedis) Set(ctx context.Context, forecastID string, forecast *Forecast) error {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if err := enc.Encode(forecast); err != nil {
		return fmt.Errorf("failed to serialize forecast structure: %s", err)
	}

	if err := cr.client.WithContext(ctx).Set(cr.keyName(forecastID), buf.String(), time.Minute*15).Err(); err != nil {
		return fmt.Errorf("failed to store forecast value: %s", err)
	}

	return nil
}
