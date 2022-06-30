package router

import (
	"MusicWebsite/controller"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func RouteInit(r *gin.Engine, indexPath string) {

	// 首页跳转路由
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = indexPath
		r.HandleContext(c)
	})

	gets(r)
	posts(r)
}

// GET 请求路由
func gets(r *gin.Engine) {
	// r.GET("/regis", controller.RegisterHandler)
	// r.GET("/log", controller.LoginHandler)

	r.GET("/songList",controller.GetSongs)
	r.GET("/getmls", controller.GetMLs)
}

// POST 请求路由
func posts(r *gin.Engine) {
	
	// 用户管理
	r.POST("/register", controller.AddUserHandler)
	r.POST("/login", controller.UserLoginHandler)
	r.POST("/updatePw", controller.UpdatePassword)
	r.POST("/updateUser", controller.UpdateUserInfo)

	// 歌单管理
	r.POST("/getml", controller.GetMusicListHandler)
	r.POST("/getUmls", controller.GetMLByUid)
	r.POST("/addStML", controller.AddSongToList)
	r.POST("/delSfML", controller.DelSongFromML)
	r.POST("/addML", controller.AddMusicList)

	// 单曲管理
	r.POST("/songId", controller.GetSongInfo)
	r.POST("/addSong", controller.AddSong)
	r.POST("/deleteSong", controller.DelSong)
	r.POST("/searchSongs", controller.SearchByKw)
	r.POST("/discover", controller.GetSongs)
	r.POST("/hot", controller.GetHotSongs)

	// 播放记录管理
	r.POST("/addRecord", controller.AddRecord)
	r.POST("/getRecordSong", controller.GetRecord)
	
	// 评论管理
	r.POST("/addCmt", controller.AddComment)
	r.POST("/getCmt", controller.GetCommentByIid)

	// 推荐算法
	r.POST("/recommendUsers", controller.GetRcmdUsers)
	r.POST("/recommendSongs", controller.GetRcmdSongs)
	r.POST("/similarSongs", controller.GetSimilarSongs)
	
}
