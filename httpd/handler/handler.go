package handler

import "GO-GIN_REST_API/cache"

type Handler struct {
	backend *cache.RedisClient
}

func NewHandler() *Handler {
	redisconf := &cache.RedisConfig{Endpoint: "localhost:6379", Password: "", Database: 0, PoolSize: 0}
	redis, err := cache.NewRedis(redisconf)
	if err != nil {
		panic(err)
	}
	return &Handler{backend: redis}
}
