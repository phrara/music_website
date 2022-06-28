package model

type Comment struct {
	Cid int `gorm:"column:cid" json:"cid"`
	Iid string `gorm:"column:iid" json:"iid"`
	Uid string `gorm:"column:uid" json:"uid"`
	Content string `gorm:"column:content" json:"content"`
	Date string `gorm:"column:date" json:"date"`
}

func (c Comment) TableName() string {
	return "comment"
}

func NewComment() *Comment {
	var cm Comment
	return &cm
}