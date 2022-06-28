package dao

import "MusicWebsite/model"

type CommentDao struct {
}

// 添加评论
func (cd *CommentDao) AddComment(c *model.Comment) bool {
	DBMgr.Create(c)
	return true
}

// 获取某歌曲的评论
func (cd *CommentDao) GetCommentByIid(iid string) []model.Comment {
	clist := make([]model.Comment,15)
	DBMgr.Where("iid = ?", iid).First(&clist)
	return clist
}

// 某用户的评论
func (cd *CommentDao) GetCommentByUid() {
}