package logic

import (
	"MicroService/kitex_gen/user"
	"MicroService/rpc/user/dao"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func Register(ctx context.Context, db *gorm.DB, rdb *redis.Client, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	exist, err := dao.CheckUsernameExist(db, req.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return &user.RegisterResponse{
			Code:    1,
			Message: "查询数据库错误",
		}, err
	}
	if exist {
		return &user.RegisterResponse{
			Code:    2,
			Message: "用户名已存在",
		}, nil
	}
	newUser := user.User{
		Username: req.Username,
		Password: req.Password,
	}
	err = dao.CreateUser(db, &newUser)
	if err != nil {
		return &user.RegisterResponse{
			Code:    3,
			Message: "注册失败",
		}, err
	}
	return &user.RegisterResponse{
		Code:    0,
		Message: "注册成功",
	}, nil
}
