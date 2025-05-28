package main

import (
	"ascue/config"
	"ascue/internal/api"
	"ascue/internal/fetch"
	"ascue/internal/redisstore"
	"ascue/internal/storage"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	rawRedis := storage.NewRedisClient(cfg.RedisAddr, cfg.RedisPass)
	store := redisstore.New(rawRedis)

	fetch.Launch(cfg.Targets, cfg.Keys, cfg.Interval, store)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", api.NewRouter(store))
}
