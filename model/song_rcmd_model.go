package model

import "strings"

type SongRCMD struct {
	Uid      string `gorm:"column:uid" json:"uid"`
	Songs string `gorm:"column:songs" json:"songs"`
}

func (s SongRCMD) TableName() string {
	return "rcmdsongs"
}

func NewSongRCMD(uid string) *SongRCMD {
	return &SongRCMD{
		Uid: uid,
		Songs: "",
	}
}

func (sr *SongRCMD) GetSongs() []string {
	return strings.Split(sr.Songs, ",")
}
