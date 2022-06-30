package model

type Song struct {
	Iid         string `gorm:"primary_key;column:iid" json:"iid"`
	SongName    string `gorm:"column:song_name" json:"song_name"`
	SongUrl     string `gorm:"column:song_url" json:"song_url"`
	SingerId    string `gorm:"column:suid" json:"suid"`
	PlayCount   uint   `gorm:"column:playcnt" json:"playcnt"`
	Album       string `gorm:"column:album" json:"album"`
	DownUrl     string `gorm:"column:down_url" json:"down_url"`
	SongTime    string `gorm:"column:song_time" json:"song_time"`
	PicUrl      string `gorm:"column:picUrl" json:"picUrl"`
	PublishTime string `gorm:"column:publishTime" json:"publishTime"`
	IsDelete    string `gorm:"column:isDelete" json:"isDelete"`
}

func (s Song) TableName() string {
	return "song"
}

// 重写Song的排序函数
// 根据播放次数进行排序
// 降序排列
type SongList []Song

func (s SongList) Len() int {
	return len(s)
}
func (s SongList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SongList) Less(i, j int) bool {
	return s[j].PlayCount < s[i].PlayCount 
}

// 构造函数
func NewSong(iid, songName string) *Song {
	return &Song{
		Iid:         iid,
		SongName:    songName,
		SongUrl:     "",
		SingerId:    "",
		PlayCount:   0,
		Album:       "",
		DownUrl:     "",
		SongTime:    "",
		PicUrl:      "",
		PublishTime: "",
		IsDelete:    "",
	}
}
