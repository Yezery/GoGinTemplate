//go:build go1.17
// +build go1.17

package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	// 创建 Gin 路由实例
	r := gin.Default()

	// 添加路由规则
	r.GET("/", func(c *gin.Context) {
		// c.String(http.StatusOK, "Hello, Gin!")
	})

	// 启动服务并监听 8080 端口
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
