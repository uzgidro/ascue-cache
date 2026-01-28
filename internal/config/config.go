package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type RedisConfig struct {
	Addr string `yaml:"addr" env:"REDIS_ADDR" env-default:"localhost:6379"`
	Pass string `yaml:"pass" env:"REDIS_PASS"`
	DB   int    `yaml:"db" env:"REDIS_DB" env-default:"0"`
}

type Config struct {
	Interval time.Duration `yaml:"interval" env:"PING_INTERVAL" env-default:"1m"`
	Targets  []string      `yaml:"targets" env:"PING_TARGETS"`
	Keys     []string      `yaml:"keys" env:"PING_KEYS"`
	Redis    RedisConfig   `yaml:"redis"`
}

func Load() Config {
	var cfg Config

	// Try to read from config file first, fallback to environment variables
	configPath := "config/config.yaml"
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Println("config.yaml not found, trying environment variables")
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Fatal("Failed to load configuration:", err)
		}
	}

	return cfg
}
