package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Enable   string `json:"enable"`
	Endpoint string `json:"endpoint"`
	Password string `json:"password"`
	Database int    `json:"database"`
	PoolSize int    `json:"poolSize"`
}

type RedisClient struct {
	Client *redis.Client
	// prefix string
}

func NewRedis(config *RedisConfig) (*RedisClient, error) {
	log.Printf("redisconf: %+v", config)
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Endpoint,
		Password: config.Password,
		DB:       config.Database,
		PoolSize: config.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("连接redis出错，错误信息：", err)
		log.Fatalf("连接redis出错，错误信息：%v", err)
		// panic(err)
		return nil, err
	}
	log.Println("redis connected successfull")

	return &RedisClient{Client: rdb}, nil
}
