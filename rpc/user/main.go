package main

import (
	"MicroService/config"
	"MicroService/consul"
	user2 "MicroService/kitex_gen/user"
	user "MicroService/kitex_gen/user/userservice"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"net"
)

type User struct {
	user2.User
	ID uint64 `gorm:"primaryKey;autoIncrement"`
}

func initDB() *gorm.DB {
	db := config.InitMysql()
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func initRedis() *redis.Client {
	rdb := config.InitRedisDB()
	return rdb
}

func main() {
	klog.SetLevel(klog.LevelDebug) // 设置日志级别为 Debug
	db := initDB()
	rdb := initRedis()

	handler := UserServiceImpl{
		db:  db,
		rdb: rdb,
	}

	// 创建 Consul 注册中心
	consulRegistry, err := consul.NewConsulRegistry("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	// 创建服务实例
	svr := user.NewServer(
		&handler,
		server.WithRegistry(consulRegistry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "UserServer",
			Tags:        map[string]string{"env": "dev"},
		}),
		server.WithServiceAddr(&net.TCPAddr{ // 设置服务地址和端口
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8888, // 你的服务端口号
		}),
	)

	// 启动服务
	if err := svr.Run(); err != nil {
		panic(err)
	}

}
