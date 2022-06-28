package model

type MusicList struct {
	Mid      int    `gorm:"primary_key;column:mid" json:"mid"`
	Uid      string `gorm:"column:uid" json:"uid"`
	ListName string `gorm:"column:listname" json:"listName"`
}

func (l MusicList) TableName() string {
	return "music_list"
}

func NewMusicList(mid int, uid, listName string) *MusicList {
	return &MusicList{
		Mid:      mid,
		Uid:      uid,
		ListName: listName,
	}
}
