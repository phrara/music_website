package controller

import (
	"MusicWebsite/model"
	"MusicWebsite/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var ss = service.NewSongService()

// @title	SearchByKw
// @description   搜索
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func SearchByKw(c *gin.Context){
	b, err := c.GetRawData()
	if err != nil {
		return
	}
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	list := ss.SearchForSongs(m["kw"].(string))
	c.JSON(200, list)
}
	

// @title	GetSongs
// @description   获取歌曲
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func GetSongs(c *gin.Context){
	list := ss.GetSongList()
	c.JSON(200, list)
}

// @title	AddSong
// @description   添加歌曲
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func AddSong(c *gin.Context){
	si := new(model.SongInfo)
	err := c.ShouldBind(si)
	if err != nil {
		return 
	}
	b := ss.AddSingleSong(si)
	c.JSON(200, b)
}

// @title	DelSongs
// @description   删除歌曲
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func DelSong(c *gin.Context){
	s := model.NewSong("","")
	err := c.ShouldBind(s)
	if err != nil {
		return 
	}
	b := ss.DelSingleSong(s)
	c.JSON(200, b)
}

// @title	GetSongInfo
// @description   获取歌曲信息
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func GetSongInfo(c *gin.Context) {
	song := new(model.Song)
	err := c.ShouldBind(song)
	if err != nil {
		return
	}
	b := ss.GetSongInfo(song)
	c.JSON(200, b)
}