package controller

import (
	"MusicWebsite/model"
	"MusicWebsite/service"
	"github.com/gin-gonic/gin"
)

// @title	AddUserHandler
// @description   添加用户（注册）
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func AddUserHandler(c *gin.Context) {
	us := service.NewUserService()
	user := model.NewUser("", "", "")
	err := c.ShouldBind(user)
	if err != nil {
		return
	}
	b := us.UserRegister(user)
	if b {
		c.JSON(200, gin.H{
			"msg":  "ok",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "err",
			"data": nil,
		})
	}
}

// @title	UserLoginHandler
// @description   登录
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func UserLoginHandler(c *gin.Context) {
	us := service.NewUserService()
	user := model.NewUser("", "", "")
	err := c.ShouldBind(user)
	if err != nil {
		return
	}
	b, user1 := us.UserLogin(user)
	if b {
		c.JSON(200, gin.H{
			"msg":  "ok",
			"data": *user1,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "err",
			"data": nil,
		})
	}
}

// @title	UpdatePassword
// @description   更改密码
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func UpdatePassword(c *gin.Context){
	us := service.NewUserService()
	info := new(model.UserInfo)
	err := c.ShouldBind(info)
	if err != nil {
		return
	}
	b := us.UpdatePassword(info)
	if b {
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





