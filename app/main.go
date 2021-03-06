package main

import (
	"MusicWebsite/router"
	"MusicWebsite/tool"
	"os/exec"
	"time"
	"github.com/gin-gonic/gin"
)


const (
	CMD = "python"
	PythonScript = "../recommend_algorithm/main_function.py"
	OutputFile = "../recommend_algorithm/output.txt"
)

// 执行推荐算法
func StartRcmdAlgo(){
	// 定时器
	// 间隔：24H
	ticker := time.Tick(time.Hour * 24)
	for i := range ticker {
		subProcess := exec.Command(CMD, PythonScript)
		b, _ := subProcess.Output()
		tool.WriteFile(OutputFile, i.Format("[2006-01-02 15:04:05]") + "\n" + string(b))
	}
	
}

func main() {
	r := gin.Default()
	// 加载 html
	//r.LoadHTMLGlob("templates/*")
	// 加载静态文件
	r.Static("/dist", "../dist")
	r.Static("/assets", "../dist/assets")
	//r.Static("/js", "./dist/js")
	router.RouteInit(r, "/dist")

	// TODO: 周期执行 python 脚本
	go StartRcmdAlgo()
	

	err := r.Run(":8082")
	if err != nil {
		return
	}
}
