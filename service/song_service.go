package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SongService struct {
	songDao *dao.SongDao
	sgrDao *dao.SingerDao
	cd *dao.CommentDao
}

func NewSongService() *SongService {
	ss := new(SongService)
	ss.songDao = new(dao.SongDao)
	ss.sgrDao = new(dao.SingerDao)
	ss.cd = new(dao.CommentDao)
	return ss
}

// 查询歌曲
func (s *SongService) SearchForSongs(keyWord string) tool.Res {
	keyWord = strings.Trim(keyWord," ")
	if keyWord == "" {
		return tool.GetGoodResult(nil)
	}
	if tool.IllegalWordsInspect(keyWord) {
		sl := s.songDao.SearchFor(keyWord)
		infoList := make([]model.SongInfo, 0)
		for _, v := range sl {
			if v.IsDelete == "0" {
				continue
			}
			sn := s.sgrDao.GetSingerInfo(v.SingerId).SingerName
			si := model.NewSongInfo(&v, sn)
			infoList = append(infoList, *si)
		}
		return tool.GetGoodResult(infoList)
	} else {
		return tool.GetBadResult("exists illegal words")
	}
}

// 获取单曲列表
func (s *SongService) GetSongList(index, size int) tool.Res {
	infoList := make([]model.SongInfo, 0)
	list := s.songDao.GetSongList(1000)
	num := 0
	for i, v := range list {
		if v.IsDelete == "0" || i < (index - 1) * size {
			continue
		}
		sn := s.sgrDao.GetSingerInfo(v.SingerId).SingerName
		si := model.NewSongInfo(&v, sn)
		infoList = append(infoList, *si)
		num++
		if num == size {
			break
		}
	}
	return tool.GetGoodResult(infoList)
}

// 未登录时获取热门歌曲
func (s *SongService) GetHotSongs(index, size int) tool.Res {
	infoList := make([]model.SongInfo, 0)
	list := s.songDao.GetSongList(1000)
	// 降序排序
	sort.Sort(model.SongList(list))
	num := 0
	for i, v := range list {
		if v.IsDelete == "0" || i < (index - 1) * size {
			continue
		}
		sn := s.sgrDao.GetSingerInfo(v.SingerId).SingerName
		si := model.NewSongInfo(&v, sn)
		infoList = append(infoList, *si)
		num++
		if num == size {
			break
		}
	}
	return tool.GetGoodResult(infoList)
}

// 获取单曲信息
func (s *SongService) GetSongInfo(song *model.Song) tool.Res {
	s2 := s.songDao.GetSongInfo(song.Iid)
	if s2.SongName != "" {
		cmt := s.cd.GetCommentByIid(song.Iid)
		data := tool.Res{
			"songInfo": *s2,
			"comment": cmt,
		}
		return tool.GetGoodResult(data)
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