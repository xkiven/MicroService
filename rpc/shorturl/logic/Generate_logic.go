package logic

import (
	"MicroService/kitex_gen/shorturl"
	"MicroService/rpc/shorturl/dao"
	"context"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

func GenerateLogic(ctx context.Context, rdb *redis.Client, req *shorturl.GenerateReq) (*shorturl.GenerateResp, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	rand.New(rand.NewSource(time.Now().UnixNano())).Intn(length + 1)
	shUrl := make([]byte, length)
	for i := range shUrl {
		shUrl[i] = charset[rand.Intn(len(charset))]
	}
	shortUrl := string(shUrl)
	longUrl := req.LongUrl

	err := dao.StoreRelation(rdb, shortUrl, longUrl)
	if err != nil {
		return nil, err
	}

	// 返回短链
	return &shorturl.GenerateResp{ShortUrl: shortUrl}, nil

}
