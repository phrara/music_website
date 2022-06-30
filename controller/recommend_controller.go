package controller

import (
	"MusicWebsite/model"
	"MusicWebsite/service"
	"MusicWebsite/tool"
	"encoding/json"

	"github.com/gin-gonic/gin"
)


var rcs = service.NewRCMDService()

// 获取推荐用户
func GetRcmdUsers(c *gin.Context) {
	userinfo := new(model.UserInfo)
	err := c.ShouldBind(userinfo)
	if err != nil {
		return
	}
	r := rcs.GetRcmdUsers(userinfo)
	c.JSON(200, r)
}


// 获取推荐歌曲
func GetRcmdSongs(c *gin.Context){
	userinfo := new(model.UserInfo)
	err := c.ShouldBind(userinfo)
	if err != nil {
		return
	}
	r := rcs.GetRcmdSongs(userinfo)
	c.JSON(200, r)
}

// 获取相似歌曲
func GetSimilarSongs(c *gin.Context){
	b, err := c.GetRawData()
	if err != nil {
		return
	}
	var data tool.Res
	_ = json.Unmarshal(b, &data)
	r := rcs.GetSimilarSongs(data["iid"].(string))
	c.JSON(200, r)
}