package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Interval  time.Duration
	Targets   []string
	Keys      []string
	RedisAddr string
	RedisPass string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using defaults")
	}

	intervalStr := os.Getenv("PING_INTERVAL")
	if intervalStr == "" {
		intervalStr = "1m"
	}
	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		interval = time.Minute
	}

	targets := strings.Split(os.Getenv("PING_TARGETS"), ",")
	keys := strings.Split(os.Getenv("PING_KEYS"), ",")

	return Config{
		Interval:  interval,
		Targets:   targets,
		Keys:      keys,
		RedisAddr: os.Getenv("REDIS_ADDR"),
		RedisPass: os.Getenv("REDIS_PASS"),
	}
}
