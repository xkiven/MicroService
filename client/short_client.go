package client

import (
	"MicroService/kitex_gen/shorturl"
	"MicroService/kitex_gen/shorturl/shorturlservice"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
)

func GenerateShortUrlClient(req *shorturl.GenerateReq) (*shorturl.GenerateResp, error) {
	// 创建 Kitex 客户端
	fmt.Println("创建客户端")
	cli, err := shorturlservice.NewClient(
		"GenerateShortUrlService",
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	resp, err := cli.Generate(context.Background(), req)
	if err != nil {
		panic(err)
	}

	return resp, nil
}

func RedirectShortUrlClient(req *shorturl.RedirectReq) (*shorturl.RedirectResp, error) {
	// 创建 Kitex 客户端
	fmt.Println("创建客户端")
	cli, err := shorturlservice.NewClient(
		"RedirectShortUrlService",
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	resp, err := cli.Redirect(context.Background(), req)
	if err != nil {
		panic(err)
	}

	return resp, nil
}
