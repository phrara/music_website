package controller

import (
	"MusicWebsite/model"
	"MusicWebsite/service"
	"github.com/gin-gonic/gin"
)

var mls = service.NewMusicListService()

// @title	GetMusicListHandler
// @description   获得歌单
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func GetMusicListHandler(c *gin.Context) {
	ml := model.NewMusicList(0, "", "")
	err := c.ShouldBind(ml)
	if err != nil {
		return
	}
	list := mls.GetMusicList(ml)
	c.JSON(200, list)
}

// @title	GetMLByUid
// @description   获得某用户歌单
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func GetMLByUid(c *gin.Context){
	ui := new(model.UserInfo)
	err := c.ShouldBind(ui)
	if err != nil {
		return
	}
	list := mls.GetOnesMusicList(ui)
	c.JSON(200, list)
}

// @title	GetMLs
// @description   获得总歌单
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func GetMLs(c *gin.Context){
	list := mls.GetMusicLists()
	c.JSON(200, list)
}

// @title	AddSongToList
// @description   添加单曲到歌单
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func AddSongToList(c *gin.Context){
	lc := model.NewListContent(0,"")
	err := c.ShouldBind(lc)
	if err != nil {
		return
	}
	b := mls.AddSong(lc)
	c.JSON(200, b)
}

// @title	DelSongFromML
// @description   删除歌单歌曲
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func DelSongFromML(c *gin.Context){
	lc := model.NewListContent(0,"")
	err := c.ShouldBind(lc)
	if err != nil {
		return
	}
	b := mls.DelSong(lc)
	c.JSON(200, b)
}

// @title	AddMusicList
// @description   建立歌单
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func AddMusicList(c *gin.Context){
	ml := model.NewMusicList(0,"","")
	err := c.ShouldBind(ml)
	if err != nil {
		return 
	}
	b := mls.AddMusicList(ml)
	c.JSON(200, b)
}