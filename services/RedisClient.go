package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/uptempotech/redistest/global"
)

// RedisClient .
type RedisClient struct {
	c *redis.Client
}

// NewRedisClient .
func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     global.RedisHost,
		Password: global.RedisPassword,
		DB:       global.DefaultDB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	return &RedisClient{
		c: client,
	}
}

// GetKey get key
func (client *RedisClient) GetKey(key string, src interface{}) error {
	val, err := client.c.Get(context.Background(), key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
	return nil
}

// SetKey set key
func (client *RedisClient) SetKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.c.Set(context.Background(), key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
