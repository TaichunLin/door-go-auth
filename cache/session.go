package cache

import (
	"GO-GIN_REST_API/entity"
	"encoding/json"
	"fmt"
	"log"
)

// func (r *RedisClient) GetAllAuth() []*entity.AuthUser {

// 	var cursor uint64
// 	var keys []string
// 	var res []*entity.AuthUser

// 	for {
// 		var err error
// 		keys, cursor, err = r.Client.Scan(ctx, cursor, "b2:dm:session:*", 0).Result()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		log.Println("keys", keys)
// 		for _, key := range keys {
// 			val, _ := r.Client.Get(ctx, key).Result()

// 			user := entity.AuthUser{}
// 			err := json.Unmarshal([]byte(val), &user)
// 			if err != nil {
// 				log.Println("rdb.GetAllAuth failed:", err)
// 			}

// 			log.Printf("%+v", &user)
// 			log.Println("&user", &user)

// 			res = append(res, &user)
// 		}

// 		if cursor == 0 {
// 			break
// 		}
// 	}

// 	return res
// }

func (r *RedisClient) SetAuth(key string, value *entity.AuthUser) error {

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
func (r *RedisClient) GetAuth(key string) *entity.AuthUser {

	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("error", err)
	}
	fmt.Println("rdb.GetAuth:", value)

	user := entity.AuthUser{}
	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		log.Println("rdb.GetAuth failed:", err)
	}
	return &user
}
