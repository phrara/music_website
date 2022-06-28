package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
	"strconv"
	"strings"
	"time"
)

type SongService struct {
	songDao *dao.SongDao
	sgrDao *dao.SingerDao
}

func NewSongService() *SongService {
	ss := new(SongService)
	ss.songDao = new(dao.SongDao)
	ss.sgrDao = new(dao.SingerDao)
	return ss
}

// 查询歌曲
func (s *SongService) SearchForSongs(keyWord string) (bool, []model.Song) {
	keyWord = strings.Trim(keyWord," ")
	if keyWord == "" {
		return false, nil
	}
	if tool.IllegalWordsInspect(keyWord) {
		return true, s.songDao.SearchFor(keyWord)
	} else {
		return false, nil
	}
}

// 获取单曲列表
func (s *SongService) GetSongList() []model.SongInfo {
	infoList := make([]model.SongInfo,1000)
	list := s.songDao.GetSongList()
	for _, v := range list {
		if v.IsDelete == "0" {
			continue
		}
		sn := s.sgrDao.GetSingerInfo(v.SingerId).SingerName
		si := model.NewSongInfo(&v, sn)
		infoList = append(infoList, *si)
	}
	return infoList
}

// 获取单曲信息
func (s *SongService) GetSongInfo(song *model.Song) (bool, *model.Song) {
	s2 := s.songDao.GetSongInfo(song.Iid)
	if s2.SongName != "" {
		return true, s2
	} else {
		return false, nil
	}
}


// 添加单曲
func (s *SongService) AddSingleSong(songInfo *model.SongInfo) (bool) {
	song := model.NewSong("","")
	song.Album = songInfo.Album
	song.DownUrl = songInfo.DownUrl
	song.PicUrl = songInfo.PicUrl
	song.SongName = songInfo.SongName
	song.SongTime =  songInfo.SongTime
	song.IsDelete = "1"
	song.PlayCount = 0
	song.SongUrl = songInfo.SongUrl
	song.PublishTime = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	song.Iid = tool.GetRandomId(song.PublishTime)
	song.SingerId = s.sgrDao.GetSuidByName(songInfo.SingerName)

	if b := s.songDao.AddSong(song); b {
		return true
	} else {
		return false
	}

}

// 删除
func (s *SongService) DelSingleSong(song *model.Song) bool {
	if b := s.songDao.DeleteSong(song.Iid); b {
		return true
	} else {
		return false
	}

}	