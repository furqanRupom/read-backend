package config

import (
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host            string
	Port            int
	Password        string
	DB              int
	UsersPendingTTL time.Duration
}

func (config RedisConfig) CreateClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:      config.Host + ":" + strconv.Itoa(config.Port),
		Password:  config.Password,
		DB:        config.DB,
		TLSConfig: nil,
	})
}
func createRedisConfig() RedisConfig {
	return RedisConfig{
		Host:     getEnvWithoutParser("REDIS_HOST", false),
		Port:     getEnv("REDIS_PORT", false, parseInt),
		Password: getEnvWithoutParser("REDIS_PASSWORD", false),
		DB:       getEnv("REDIS_DB", false, parseInt),
	}
}

