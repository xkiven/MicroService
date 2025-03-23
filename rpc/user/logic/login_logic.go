package logic

import (
	"MicroService/kitex_gen/user"
	"MicroService/rpc/user/dao"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
)

func Login(ctx context.Context, db *gorm.DB, rdb *redis.Client, req *user.LoginRequest) (*user.LoginResponse, error) {
	//先在redis中查找
	isLoggedIn, err := dao.IsUserLoginState(rdb, req.Username)
	if err != nil {
		return &user.LoginResponse{
			Code:    4,
			Message: "redis查询出错",
		}, err
	}
	if isLoggedIn {
		//生成JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": req.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString([]byte("your_secret_key"))
		if err != nil {
			return nil, err
		}

		return &user.LoginResponse{
			Code:    0,
			Message: "登录成功",
			Token:   tokenString,
		}, nil
	}

	userModel, err := dao.GetUserByPasswordAndUsername(db, req.Password, req.Username)
	if err != nil {
		return &user.LoginResponse{
			Code:    5,
			Message: "密码或用户名错误",
		}, err
	}
	err = dao.SetUserLoginState(rdb, userModel.Username)
	if err != nil {
		return &user.LoginResponse{
			Code:    6,
			Message: "登录缓存设置失败",
		}, err
	}

	//生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return nil, err
	}

	return &user.LoginResponse{
		Code:    0,
		Message: "登录成功",
		Token:   tokenString,
	}, nil
}
