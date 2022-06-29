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
	r.Static("/dist", "./dist")
	r.Static("/assets", "./dist/assets")
	//r.Static("/js", "./dist/js")
	router.RouteInit(r, "/dist")

	err := r.Run(":8082")
	if err != nil {
		return
	}
}
