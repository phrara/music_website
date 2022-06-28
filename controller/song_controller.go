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
	f, list := ss.SearchForSongs(m["kw"].(string))
	if f {
		c.JSON(200, gin.H{
			"msg": "ok",
			"data": list,
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "err",
			"data": list,
		})
	}
}
	

// @title	GetSongs
// @description   获取歌曲
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func GetSongs(c *gin.Context){
	list := ss.GetSongList()
	c.JSON(200, gin.H{
		"msg": "ok",
		"data": list,
	})
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
	if b := ss.AddSingleSong(si); b {
		c.JSON(200, gin.H{
			"msg": "ok",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "err",
			"data": nil,
		})
	}

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
	if b := ss.DelSingleSong(s); b {
		c.JSON(200, gin.H{
			"msg": "ok",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "err",
			"data": nil,
		})
	}

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
	b, s := ss.GetSongInfo(song)
	if b {
		c.JSON(200, gin.H{
			"msg": "ok",
			"data": *s,
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "err",
			"data": nil,
		})
	}
}