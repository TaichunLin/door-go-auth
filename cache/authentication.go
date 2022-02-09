package cache

import (
	"GO-GIN_REST_API/entity"
	"context"
	"encoding/json"
	"log"
)

func (r *RedisClient) FetchAuthen(key string) *entity.Accounts {
	ctx := context.Background()
	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("error", err)
	}
	user := entity.Accounts{}
	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		log.Println("rdb.FetchAuthen failed:", err)
	}
	return &user
}
