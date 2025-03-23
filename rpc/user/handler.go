package main

import (
	user "MicroService/kitex_gen/user"
	"MicroService/rpc/user/logic"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserServiceImpl(db *gorm.DB, rdb *redis.Client) *UserServiceImpl {
	return &UserServiceImpl{
		db:  db,
		rdb: rdb,
	}
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp, err = logic.Register(ctx, s.db, s.rdb, req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp, err = logic.Login(ctx, s.db, s.rdb, req)
	return resp, err
}
