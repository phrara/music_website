package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
	"time"
)

type CommentService struct {
	cd *dao.CommentDao
}

func NewCommentService() *CommentService{
	cs := new(CommentService)
	cs.cd = new(dao.CommentDao)
	return cs
}


func (cs *CommentService) AddComment(c *model.Comment) tool.Res {
	if tool.IllegalWordsInspect(c.Content) {
		c.Date = time.Now().Format("2006-01-02 15:04:05")
		cs.cd.AddComment(c)
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("faild")
	}
}

func (cs *CommentService) GetCommentByIid(c *model.Comment) tool.Res {
	c2 := cs.cd.GetCommentByIid(c.Iid)
	return tool.GetGoodResult(c2)
}