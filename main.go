package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/statics", "/Users/luliang/GoLand/pacman/statics")
	r.LoadHTMLGlob("/Users/luliang/GoLand/pacman/statics/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":9092")
}
