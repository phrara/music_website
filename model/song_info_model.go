package model

/*

	iid,
    song_name,
   	singer_name,
   	album,
	playcnt,
  	song_url,
   	down_url,
   	song_time,
   	picUrl,
   	publishTime

*/

type SongInfo struct {
	Iid string
	SongName string
	SingerName string
	Album string
	Playcnt uint
	SongUrl string
	DownUrl string
	SongTime string
	PicUrl string
	PublishTime string
}

func NewSongInfo(s *Song, singerName string) *SongInfo {
	return &SongInfo{
		Iid: s.Iid,
		SongName: s.SongName,
		SingerName: singerName,
		Album: s.Album,
		Playcnt: s.PlayCount,
		SongUrl: s.SongUrl,
		DownUrl: s.DownUrl,
		SongTime: s.SongTime,
		PicUrl: s.PicUrl,
		PublishTime: s.PublishTime,
	}
}