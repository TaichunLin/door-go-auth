package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"GO-GIN_REST_API/entity"
)

var ctx = context.Background()

func (r *RedisClient) GetAllGroup() []*entity.Group {

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
func (r *RedisClient) GetAllUsers() []*entity.UserList {

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

func (r *RedisClient) SetUser(key string, value *entity.User) error {

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
func (r *RedisClient) SetGroup(key string, value *entity.Group) error {

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

	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("error", err)
	}
	fmt.Println("GetGroup:", value)

	group := entity.Group{}
	err = json.Unmarshal([]byte(value), &group)
	if err != nil {
		log.Println("GetGroup failed:", err)
	}
	return &group
}

func (r *RedisClient) GetUser(key string) *entity.UserList {

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

func (r *RedisClient) Del(key string) error {
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		log.Println("Del failed:", err)
	}
	return err
}
