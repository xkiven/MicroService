package main

import (
	"MicroService/http_handler"
	"github.com/gin-gonic/gin"
)

func main() {

	// 创建 Gin 实例
	r := gin.Default()

	// 注册路由
	r.POST("/register", http_handler.RegisterHandler)
	r.POST("/login", http_handler.LoginHandler)

	r.POST("/create", http_handler.GenerateShortUrlHandler) // 创建短链
	r.GET("/:shortUrl", http_handler.RegisterHandler)       // 访问短链

	// 启动 HTTP 服务

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
