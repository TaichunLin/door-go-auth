package cache

import (
	"GO-GIN_REST_API/auth"
	"encoding/json"
	"fmt"
	"log"
)

func (r *RedisClient) SetAuth(key string, value *auth.User) error {

	byteSlice, _ := json.Marshal(value)
	log.Println(byteSlice)
	json := string(byteSlice)
	log.Println(json)
	err := r.Client.Set(ctx, key, byteSlice, 0).Err()
	if err != nil {
		log.Println("rdb.SetAuth failed:", err)
	}
	return err
}
func (r *RedisClient) GetAuth(key string) *auth.User {

	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("error", err)
	}
	fmt.Println("rdb.GetAuth:", value)

	user := auth.User{}
	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		log.Println("rdb.GetAuth failed:", err)
	}
	return &user
}
