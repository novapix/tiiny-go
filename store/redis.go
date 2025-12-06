package store

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore(redisURL string) *RedisStore {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		panic("Invalid Redis URL: " + err.Error())
	}

	client := redis.NewClient(options)
	return &RedisStore{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisStore) Save(key, url string) error {
	err := r.client.Set(r.ctx, key, url, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisStore) Get(key string) (string, error) {
	result, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("key not found")
	}
	return result, err
}
