package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"GO-GIN_REST_API/entity"
)

//Del() in cacheGroup.go

func (r *RedisClient) GetAllGroup() []*entity.Group {
	ctx := context.Background()
	var cursor uint64
	var keys []string
	var res []*entity.Group

	for {
		var err error
		keys, cursor, err = r.Client.Scan(ctx, cursor, "b2:dm:group:*", 0).Result()
		if err != nil {
			log.Println(err)
		}
		log.Println("keys", keys)
		for _, key := range keys {
			val, _ := r.Client.Get(ctx, key).Result()

			group := entity.Group{}
			err := json.Unmarshal([]byte(val), &group)
			if err != nil {
				log.Println("rdb.Get failed:", err)
			}

			log.Printf("%+v", &group)
			log.Println("&group", &group)

			res = append(res, &group)
		}

		if cursor == 0 {
			break
		}
	}

	return res
}

func (r *RedisClient) SetUser(key string, value *entity.User) error {

	ctx := context.Background()
	byteSlice, _ := json.Marshal(value)
	log.Println(byteSlice)
	json := string(byteSlice)
	log.Println(json)
	err := r.Client.Set(ctx, key, byteSlice, 0).Err()
	if err != nil {
		log.Println("rdb.Set failed:", err)
	}
	return err
}

func (r *RedisClient) GetUser(key string) *entity.UserList {
	ctx := context.Background()
	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("error", err)
	}
	fmt.Println("GetUser:", value)

	user := entity.User{}
	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		log.Println("GetUser failed:", err)
	}

	ul := &entity.UserList{Username: user.Username, CardId: user.CardId, GroupId: user.GroupId.GroupId, Group: user.GroupId.Group}
	log.Print(ul)
	return ul
}
