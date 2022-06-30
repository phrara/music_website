package controller

import (
	"MusicWebsite/model"
	"MusicWebsite/service"

	"github.com/gin-gonic/gin"
)


var urs = service.NewUserRCMDService()

// 获取推荐用户
func GetRcmdUsers(c *gin.Context) {
	userinfo := new(model.UserInfo)
	err := c.ShouldBind(userinfo)
	if err != nil {
		return
	}
	r := urs.GetRcmdUsers(userinfo)
	c.JSON(200, r)
}


// 获取推荐歌曲
func GetRcmdSongs(c *gin.Context){
	
}