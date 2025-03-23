package dao

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func SetUserLoginState(rdb *redis.Client, username string) error {
	return rdb.Set(ctx, "user:"+username, "logged_in", 0).Err()
}

func IsUserLoginState(rdb *redis.Client, username string) (bool, error) {
	value, err := rdb.Get(ctx, "user:"+username).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return value == "logged_in", nil
}
