package router

import (
	"MusicWebsite/controller"
	"github.com/gin-gonic/gin"
)

func RouteInit(r *gin.Engine, indexPath string) {

	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = indexPath
		r.HandleContext(c)
	})

	gets(r)
	posts(r)
}

func gets(r *gin.Engine) {
	// r.GET("/regis", controller.RegisterHandler)
	// r.GET("/log", controller.LoginHandler)

	r.GET("/songList",controller.GetSongs)
	r.GET("/getmls", controller.GetMLs)
}

func posts(r *gin.Engine) {
	
	r.POST("/register", controller.AddUserHandler)
	r.POST("/login", controller.UserLoginHandler)
	r.POST("/updatePw", controller.UpdatePassword)

	r.POST("/getml", controller.GetMusicListHandler)
	r.POST("/getUmls", controller.GetMLByUid)
	r.POST("/addStML", controller.AddSongToList)
	r.POST("/delSfML", controller.DelSongFromML)
	r.POST("/addML", controller.AddMusicList)

	r.POST("/songId", controller.GetSongInfo)
	r.POST("/addSong", controller.AddSong)
	r.POST("/deleteSong", controller.DelSong)
	r.POST("/search", controller.SearchByKw)

	r.POST("/addRecord", controller.AddRecord)
	r.POST("/getRecordSong", controller.GetRecord)
	
	r.POST("/addCmt", controller.AddComment)
	r.POST("/getCmt", controller.GetCommentByIid)
	
}
