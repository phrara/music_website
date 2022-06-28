package model

type ListContent struct {
	Id  int    `gorm:"primary_key;column:id" json:"id"`
	Mid int    `gorm:"column:mid" json:"mid"`
	Iid string `gorm:"column:iid" json:"iid"`
}

func (l ListContent) TableName() string {
	return "list_content"
}

func NewListContent(mid int, iid string) *ListContent {
	return &ListContent{
		Id:  0,
		Mid: mid,
		Iid: iid,
	}
}
