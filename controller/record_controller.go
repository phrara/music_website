package controller

import (
	"MusicWebsite/model"
	"MusicWebsite/service"

	"github.com/gin-gonic/gin"
)

// @title	GetRecord
// @description   获取记录
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func GetRecord(c *gin.Context){
	rs := service.NewRecordService()
	r := model.NewRecord("","")
	err := c.ShouldBind(r)
	if err != nil {
		return
	}
	rec := rs.GetRecord(r)
	c.JSON(200, gin.H{
		"msg": "ok",
		"data": rec,
	})
}

// @title	AddRecord
// @description   建立（添加）记录
// @auth      彭竑睿           时间（2022/6/28）
// @param      c      *gin.Context       "HttpContent"
// @return    无       无      "无"
func AddRecord(c *gin.Context){
	rs := service.NewRecordService()
	r := model.NewRecord("","")
	err := c.ShouldBind(r)
	if err != nil {
		return
	}
	rec := rs.AddRecord(r)
	c.JSON(200, gin.H{
		"msg": "ok",
		"data": rec,
	})
}