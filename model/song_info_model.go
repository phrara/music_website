package model

import "MusicWebsite/tool"

// 音乐种类
const (
	ROCK = iota
	RB
	JAZZ
	CLASSIC
	POP
	RAP
	FOLK
	PURE
	ELECTRONIC
	COUNTRYSIDE
)

type SongInfo struct {
	Iid string `json:"iid"`
	SongName string `json:"songName"`
	SingerName string `json:"singerName"`
	Album string `json:"album"`
	Playcnt uint `json:"playcnt"`
	SongUrl string `json:"songUrl"`
	DownUrl string `json:"downUrl"`
	SongTime string `json:"songTime"`
	PicUrl string `json:"picUrl"`
	PublishTime string `json:"publishTime"`
}

func NewSongInfo(s *Song, singerName string) *SongInfo {
	// 歌曲时长格式转换
	duration := tool.DurationTransform(s.SongTime)
	return &SongInfo{
		Iid: s.Iid,
		SongName: s.SongName,
		SingerName: singerName,
		Album: s.Album,
		Playcnt: s.PlayCount,
		SongUrl: s.SongUrl,
		DownUrl: s.DownUrl,
		SongTime: duration,
		PicUrl: s.PicUrl,
		PublishTime: s.PublishTime,
	}
}