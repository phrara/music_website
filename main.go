package main

import (
	"MusicWebsite/router"
	"github.com/gin-gonic/gin"
)


const (
	CMD = "python3"
	PythonScript = "./recommend_algorithm/main_function.py"
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

	// TODO: 周期执行 python 脚本
	//exec.Command(CMD, PythonScript)

	err := r.Run(":8082")
	if err != nil {
		return
	}
}
