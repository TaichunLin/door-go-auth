package cache

import (
	"GO-GIN_REST_API/entity"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (r *RedisClient) CreateAuthor(account string, tm *entity.TokenMetadata) error {
	ctx := context.Background()

	at := time.Unix(tm.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(tm.RtExpires, 0)
	now := time.Now()

	AccessKey := "jwtMetadata:AccessUuid:" + tm.AccessUuid
	val := r.FetchAuthen(`b2:dm:account:` + account)
	log.Println("val: ", val)
	byteSlice, _ := json.Marshal(val)
	log.Println(byteSlice)
	json := string(byteSlice)
	log.Println("json: ", json)
	errAccess := r.Client.Set(ctx, AccessKey, byteSlice, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	Refreshkey := "jwtMetadata:RefreshUuid:" + tm.RefreshUuid

	errRefresh := r.Client.Set(ctx, Refreshkey, byteSlice, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

//check if the metadata, which is from ExtractTokenMetadata(), still exists in our Redis store.
func (r *RedisClient) FetchAuthor(key *entity.AccessDetails) (string, error) {
	ctx := context.Background()

	val, err := r.Client.Get(ctx, `jwtMetadata:AccessUuid:`+key.AccessUuid).Result()
	if err != nil {
		return "", err
	}
	log.Println("FetchAuthor:")
	log.Println(val)

	return val, nil
}

func (r *RedisClient) DeleteAuthor(givenUuid string) (int64, error) {
	ctx := context.Background()

	deleted, err := r.Client.Del(ctx, givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
