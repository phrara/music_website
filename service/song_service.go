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
func (s *SongService) SearchForSongs(keyWord string) tool.Res {
	keyWord = strings.Trim(keyWord," ")
	if keyWord == "" {
		return tool.GetGoodResult(nil)
	}
	if tool.IllegalWordsInspect(keyWord) {
		return tool.GetGoodResult(s.songDao.SearchFor(keyWord))
	} else {
		return tool.GetBadResult("exists illegal words")
	}
}

// 获取单曲列表
func (s *SongService) GetSongList() tool.Res {
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
	return tool.GetGoodResult(infoList)
}

// 获取单曲信息
func (s *SongService) GetSongInfo(song *model.Song) tool.Res {
	s2 := s.songDao.GetSongInfo(song.Iid)
	if s2.SongName != "" {
		return tool.GetGoodResult(*s2)
	} else {
		return tool.GetBadResult("failed")
	}
}


// 添加单曲
func (s *SongService) AddSingleSong(songInfo *model.SongInfo) tool.Res {
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
	song.Iid = tool.GetRandomId(song.PublishTime + song.SongName)
	song.SingerId = s.sgrDao.GetSuidByName(songInfo.SingerName)

	if b := s.songDao.AddSong(song); b {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("failed")
	}

}

// 删除
func (s *SongService) DelSingleSong(song *model.Song) tool.Res {
	if b := s.songDao.DeleteSong(song.Iid); b {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("failed")
	}

}	