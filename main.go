package main

import (
	"ascue/internal/api"
	"ascue/internal/config"
	"ascue/internal/fetch"
	"ascue/internal/redisstore"
	"ascue/internal/storage"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.Load()

	rawRedis := storage.NewRedisClient(cfg.Redis.Addr, cfg.Redis.Pass, cfg.Redis.DB)
	store := redisstore.New(rawRedis)

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	fetch.Launch(cfg.Targets, cfg.Keys, cfg.Interval, store, httpClient)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", api.NewRouter(store))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
