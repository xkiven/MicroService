package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func StoreRelation(rdb *redis.Client, shortUrl string, longUrl string) error {
	// 存储短链和长 URL 的映射关系
	err := rdb.Set(ctx, shortUrl, longUrl, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func SearchLongUrl(rdb *redis.Client, shortUrl string) (string, error) {
	longUrl, err := rdb.Get(ctx, shortUrl).Result()
	if err != nil {
		return "", err

	}
	return longUrl, nil
}
