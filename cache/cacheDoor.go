package cache

import (
	"context"
	"encoding/json"
	"log"

	"GO-GIN_REST_API/entity"
)

//cacheDoor.go

type Todb interface {
	GetAllGroup() []*entity.Group
	GetAllUsers() []*entity.UserList
	Set(key string, value interface{}) error
	GetGroup(key string) *entity.Group
	GetUser(key string) *entity.UserList
	Del(key string) error
}

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
		for _, key := range keys {
			val, _ := r.Client.Get(ctx, key).Result()

			group := entity.Group{}
			err := json.Unmarshal([]byte(val), &group)
			if err != nil {
				log.Println("rdb.Get failed:", err)
			}
			res = append(res, &group)
		}

		if cursor == 0 {
			break
		}
	}

	return res
}

func (r *RedisClient) GetUser(key string) *entity.UserList {
	ctx := context.Background()
	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("error", err)
	}
	user := entity.User{}
	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		log.Println("GetUser failed:", err)
	}

	ul := &entity.UserList{Username: user.Username, CardId: user.CardId, GroupId: user.GroupId.GroupId, Group: user.GroupId.Group}
	log.Print(ul)
	return ul
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
		for _, key := range keys {
			val, _ := r.Client.Get(ctx, key).Result()

			user := entity.User{}
			err := json.Unmarshal([]byte(val), &user)
			if err != nil {
				log.Println("GetAllUsers():", err)
			}
			res = append(res, &user)
		}
		if cursor == 0 {
			break
		}
	}
	ul := []*entity.UserList{}
	for _, user := range res {
		ul = append(ul, &entity.UserList{Username: user.Username, CardId: user.CardId, GroupId: user.GroupId.GroupId, Group: user.GroupId.Group})
	}
	return ul
}

func (r *RedisClient) Set(key string, value interface{}) error {

	ctx := context.Background()
	byteSlice, _ := json.Marshal(value)
	err := r.Client.Set(ctx, key, byteSlice, 0).Err()
	if err != nil {
		log.Println(value)
		log.Println("rdb.Set failed:", err)
	}
	return err
}

func (r *RedisClient) Del(key string) error {
	ctx := context.Background()
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		log.Println("Del failed:", err)
	}
	return err
}
