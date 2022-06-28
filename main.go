package main

import (
	"MusicWebsite/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 加载 html
	//r.LoadHTMLGlob("templates/*")
	// 加载静态文件
	r.Static("/q", "./dist")
	r.Static("/css", "./dist/css")
	r.Static("/js", "./dist/js")
	router.RouteInit(r, "/q")

	err := r.Run(":8081")
	if err != nil {
		return
	}
}
