package redisstore

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Store interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type RedisStore struct {
	client *redis.Client
}

func New(client *redis.Client) Store {
	return &RedisStore{client: client}
}

func (r *RedisStore) Set(key string, value []byte) error {
	return r.client.Set(context.Background(), key, value, 0).Err()
}

func (r *RedisStore) Get(key string) ([]byte, error) {
	val, err := r.client.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}
