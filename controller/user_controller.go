package controller

import (
	"MusicWebsite/model"
	"MusicWebsite/service"
	"github.com/gin-gonic/gin"
)

var us = service.NewUserService()

// @title	AddUserHandler
// @description   添加用户（注册）
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func AddUserHandler(c *gin.Context) {
	user := model.NewUser("", "", "")
	err := c.ShouldBind(user)
	if err != nil {
		return
	}
	res := us.UserRegister(user)
	c.JSON(200, res)
}

// @title	UserLoginHandler
// @description   登录
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func UserLoginHandler(c *gin.Context) {
	user := model.NewUser("", "", "")
	err := c.ShouldBind(user)
	if err != nil {
		return
	}
	b := us.UserLogin(user)
	c.JSON(200, b)
}

// @title	UpdatePassword
// @description   更改密码
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func UpdatePassword(c *gin.Context){
	info := new(model.UserInfo)
	err := c.ShouldBind(info)
	if err != nil {
		return
	}
	b := us.UpdatePassword(info)
	c.JSON(200, b)
}

// 用户注销
func DelUser(c *gin.Context){
	info := new(model.UserInfo)
	err := c.ShouldBind(info)
	if err != nil {
		return
	}
	b := us.DeleteUser(info)
	c.JSON(200, b)
}


// 更改用户信息
func UpdateUserInfo(c *gin.Context){
	user := model.NewUser("", "", "")
	err := c.ShouldBind(user)
	if err != nil {
		return
	}
	res := us.UpdateInfo(user)
	c.JSON(200, res)
}

// 获取所有用户
func GetAllUsers(c *gin.Context) {
	r := us.GetAllUsers()
	c.JSON(200, r)
}
