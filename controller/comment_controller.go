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
	if b := cs.AddComment(cmt); b {
		c.JSON(200, gin.H{
			"msg": "ok",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "illegalWords",
			"data": nil,
		})
	}
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
	clist := cs.GetCommentByIid(cmt)
	c.JSON(200, gin.H{
			"msg": "ok",
			"data": clist,
		})
}