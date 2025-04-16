package storage

import "github.com/redis/go-redis/v9"

func NewRedisClient(addr string, pass string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})
}
