package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/xxx", "/root/pacman/statics")
	r.LoadHTMLGlob("/root/pacman/index.html")


	r.GET("/", func(c *gin.Context) {
		http.Post("http://39.104.17.28:9001/Notice", "text/plain",
			strings.NewReader("又有用户访问你的网站啦~😀"))
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":80")
}
