package client

import (
	"MicroService/kitex_gen/user"
	"MicroService/kitex_gen/user/userservice"
	"context"
	"github.com/cloudwego/kitex/client"
	"time"
)

// RegisterClientWithRequest 接收请求参数并调用 Kitex 服务
func RegisterClientWithRequest(req *user.RegisterRequest) (*user.RegisterResponse, error) {
	cli, err := userservice.NewClient(
		"UserRegisterService",
		client.WithHostPorts("127.0.0.1:8888"), // 服务地址
	)
	if err != nil {
		return nil, err
	}

	// 调用远程方法
	resp, err := cli.Register(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// LoginClientWithRequest 接收请求参数并调用 Kitex 服务
func LoginClientWithRequest(req *user.LoginRequest) (*user.LoginResponse, error) {
	cli, err := userservice.NewClient(
		"UserLoginService",
		client.WithHostPorts("127.0.0.1:8888"), // 服务地址
		client.WithRPCTimeout(5*time.Second),   // 设置超时时间
	)
	if err != nil {
		return nil, err
	}

	// 调用远程方法
	resp, err := cli.Login(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
