package controller

import (
	"MusicWebsite/model"
	"MusicWebsite/service"
	"github.com/gin-gonic/gin"
)

var cs = service.NewCommentService()

// @title	AddComment
// @description   添加评论
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func AddComment(c *gin.Context){
	cmt := model.NewComment()
	err := c.ShouldBind(cmt)
	if err != nil {
		return
	}
	b := cs.AddComment(cmt)
	c.JSON(200, b)
}

// @title	GetCommentByIid
// @description   获取评论
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func GetCommentByIid(c *gin.Context){
	cmt := model.NewComment()
	err := c.ShouldBind(cmt)
	if err != nil {
		return
	}
	res := cs.GetCommentByIid(cmt)
	c.JSON(200, res)
}