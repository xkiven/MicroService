package logic

import (
	"MicroService/kitex_gen/shorturl"
	"MicroService/rpc/shorturl/dao"
	"context"
	"github.com/go-redis/redis/v8"
)

func RedirectLogic(ctx context.Context, rdb *redis.Client, req *shorturl.RedirectReq) (*shorturl.RedirectResp, error) {
	// 查询长 URL
	longUrl, err := dao.SearchLongUrl(rdb, req.ShortUrl)
	if err != nil {
		return nil, err
	}

	// 返回长 URL
	return &shorturl.RedirectResp{LongUrl: longUrl}, nil

}
