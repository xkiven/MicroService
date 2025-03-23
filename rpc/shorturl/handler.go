package main

import (
	shorturl "MicroService/kitex_gen/shorturl"
	"MicroService/rpc/shorturl/logic"
	"context"
	"github.com/go-redis/redis/v8"
)

// ShortUrlServiceImpl implements the last service interface defined in the IDL.
type ShortUrlServiceImpl struct {
	rdb *redis.Client
}

// Generate implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) Generate(ctx context.Context, req *shorturl.GenerateReq) (resp *shorturl.GenerateResp, err error) {
	resp, err = logic.GenerateLogic(ctx, s.rdb, req)

	return resp, err
}

// Redirect implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) Redirect(ctx context.Context, req *shorturl.RedirectReq) (resp *shorturl.RedirectResp, err error) {
	resp, err = logic.RedirectLogic(ctx, s.rdb, req)

	return resp, err
}
