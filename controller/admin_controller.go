package controller

import (
	"MusicWebsite/model"
	"MusicWebsite/service"
	"github.com/gin-gonic/gin"
)

var as = service.NewAdminService()


// @title	AdminLogin
// @description   管理员登录
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func AdminLogin(c *gin.Context){
	admin := model.NewAdmin("","")
	err := c.ShouldBind(admin)
	if err != nil {
		return
	}

	b := as.AdminLogin(admin)
	c.JSON(200, b)
}