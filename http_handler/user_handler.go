package http_handler

import (
	"MicroService/client"
	"MicroService/kitex_gen/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterHandler 处理注册请求
func RegisterHandler(c *gin.Context) {
	var req user.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"data":  1000,
		})
		return
	}

	// 调用 Kitex 客户端
	resp, err := client.RegisterClientWithRequest(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  1001,
		})
		return
	}

	// 返回 HTTP 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "Register successful",
		"data":    resp,
	})
}

// LoginHandler 处理登录请求
func LoginHandler(c *gin.Context) {
	var req user.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"data":  1000,
		})
		return
	}

	// 调用 Kitex 客户端
	resp, err := client.LoginClientWithRequest(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  1002,
		})
		return
	}

	// 返回 HTTP 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data":    resp,
	})
}
