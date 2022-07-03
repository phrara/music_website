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
	// 总歌曲列表
	r.GET("/songList",controller.GetSongs)
	// 总歌单列表
	r.GET("/allMLs", controller.GetMLs)
	// 总用户列表
	r.GET("/allUsers", controller.GetAllUsers)
	// 数据统计
	r.GET("/statistics", controller.GetStatistics)
}

// POST 请求路由
func posts(r *gin.Engine) {
	
	// 后台
	r.POST("/adminLogin", controller.AdminLogin)

	// 用户管理
	r.POST("/register", controller.AddUserHandler)
	r.POST("/login", controller.UserLoginHandler)
	r.POST("/updatePw", controller.UpdatePassword)
	r.POST("/updateUser", controller.UpdateUserInfo)
	r.POST("/delUser", controller.DelUser)

	// 歌单管理
	// 获取某歌单中单曲的列表 mid
	r.POST("/getml", controller.GetMusicListHandler)
	// 删除 mid
	r.POST("/delml", controller.DeleteMuisicList)
	// 获取歌单列表 uid
	r.POST("/getUmls", controller.GetMLByUid)
	// 向某歌单添加某歌曲 mid, iid
	r.POST("/addStML", controller.AddSongToList)
	// 从某歌单删除某歌曲 mid, iid
	r.POST("/delSfML", controller.DelSongFromML)
	// 添加歌单 uid, listname
	r.POST("/addMl", controller.AddMusicList)
	

	// 单曲管理
	r.POST("/songId", controller.GetSongInfo)
	r.POST("/addSong", controller.AddSong)
	r.POST("/deleteSong", controller.DelSong)
	r.POST("/searchSongs", controller.SearchByKw)
	r.POST("/discover", controller.GetSongs)
	r.POST("/display", controller.GetSongs)
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
