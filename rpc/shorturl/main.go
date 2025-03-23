package main

import (
	"MicroService/config"
	"MicroService/consul"
	shorturl "MicroService/kitex_gen/shorturl/shorturlservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/go-redis/redis/v8"
	"net"
)

func initRedis() *redis.Client {
	rdb := config.InitRedisDB()
	return rdb
}

func main() {

	rdb := initRedis()
	handler := ShortUrlServiceImpl{
		rdb: rdb,
	}

	// 创建 Consul 注册中心
	consulRegistry, err := consul.NewConsulRegistry("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	svr := shorturl.NewServer(
		&handler,
		server.WithRegistry(consulRegistry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "ShortUrlServer",
			Tags:        map[string]string{"env": "dev"},
		}),
		server.WithServiceAddr(&net.TCPAddr{ // 设置服务地址和端口
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8889, // 你的服务端口号
		}),
	)

	// 启动服务
	if err := svr.Run(); err != nil {
		panic(err)
	}
}
