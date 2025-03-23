package http_handler

import (
	"MicroService/client"
	"MicroService/kitex_gen/shorturl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateShortUrlHandler(c *gin.Context) {
	var req shorturl.GenerateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"data":  1000,
		})
		return
	}

	// 调用 Kitex 客户端
	resp, err := client.GenerateShortUrlClient(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  1001,
		})
		return
	}

	// 返回 HTTP 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "GenerateShortUrl successful",
		"data":    resp,
	})

}

func RedirectShortUrlHandler(c *gin.Context) {
	var req shorturl.RedirectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"data":  1000,
		})
		return
	}

	// 调用 Kitex 客户端
	resp, err := client.RedirectShortUrlClient(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  1001,
		})
		return
	}

	// 返回 HTTP 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "RedirectShortUrl successful",
		"data":    resp,
	})
}
