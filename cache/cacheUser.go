package cache

import (
	"GO-GIN_REST_API/entity"
	"context"
	"encoding/json"
	"log"
)

func (r *RedisClient) GetAllUsers() []*entity.UserList {
	ctx := context.Background()
	var cursor uint64
	var keys []string
	var res []*entity.User

	for {
		var err error
		keys, cursor, err = r.Client.Scan(ctx, cursor, "b2:dm:user:*", 0).Result()
		if err != nil {
			panic(err)
		}
		log.Println("keys:", keys)

		for _, key := range keys {
			val, _ := r.Client.Get(ctx, key).Result()

			user := &entity.User{}
			err := json.Unmarshal([]byte(val), &user)
			if err != nil {
				log.Println("GetAllUsers():", err)
			}

			log.Printf("%+v", &user)
			log.Println("&user:", &user)
			res = append(res, user)
		}
		if cursor == 0 {
			break
		}
	}
	log.Println("keys2222:", keys)

	ul := []*entity.UserList{}
	for _, user := range res {
		ul = append(ul, &entity.UserList{Username: user.Username, CardId: user.CardId, GroupId: user.GroupId.GroupId, Group: user.GroupId.Group})
	}
	log.Println("ul:", ul)
	return ul
}

func (r *RedisClient) SetGroup(key string, value *entity.Group) error {
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
func (r *RedisClient) GetGroup(key string) *entity.Group {
	ctx := context.Background()
	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("error", err)
	}
	group := entity.Group{}
	err = json.Unmarshal([]byte(value), &group)
	if err != nil {
		log.Println("GetGroup failed:", err)
	}
	return &group
}

func (r *RedisClient) Del(key string) error {
	ctx := context.Background()
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		log.Println("Del failed:", err)
	}
	return err
}
