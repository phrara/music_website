package model

import "strings"

type SongSimilar struct {
	Iid      string `gorm:"column:iid" json:"iid"`
	TopSongs string `gorm:"column:topsongs" json:"topsongs"`
}

func (s SongSimilar) TableName() string {
	return "topsongs"
}

func NewSongSimilar(iid string) *SongSimilar {
	return &SongSimilar{
		Iid: iid,
		TopSongs: "",
	}
}

func (ss *SongSimilar) GetTopSongs() []string {
	return strings.Split(ss.TopSongs, ",")
}
