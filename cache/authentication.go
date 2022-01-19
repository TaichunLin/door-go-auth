package cache

import (
	"GO-GIN_REST_API/entity"
	"context"
	"encoding/json"
	"log"
)

func (r *RedisClient) CreateAuthen(key string, value *entity.Accounts) error {

	ctx := context.Background()
	byteSlice, _ := json.Marshal(value)
	err := r.Client.Set(ctx, key, byteSlice, 0).Err()
	if err != nil {
		log.Println("rdb.CreateAuthen failed:", err)
	}
	return err
}

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

//需要刪除會員嗎...
// func (r *RedisClient) DeleteAuthen(key string) error {
// 	ctx := context.Background()
// 	err := r.Client.Del(ctx, key).Err()
// 	if err != nil {
// 		log.Println("Del failed:", err)
// 	}
// 	return err
// }
