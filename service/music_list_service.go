package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
)

type MusicListService struct {
	mld *dao.MusicListDao
	songDao *dao.SongDao
}

func NewMusicListService() *MusicListService {
	mls := new(MusicListService)
	mls.mld = new(dao.MusicListDao)
	return mls
}

// 获取某用户总歌单
func (s *MusicListService) GetOnesMusicList(userInfo *model.UserInfo) tool.Res {
	list := s.mld.GetMusicListsByUid(userInfo.Uid)
	return tool.GetGoodResult(list)
}

// 获取总歌单
func (m *MusicListService) GetMusicLists() tool.Res {
	return tool.GetGoodResult(m.mld.GetMusicLists()) 
}

// GetMusicList 获取歌单
func (s *MusicListService) GetMusicList(ml *model.MusicList) tool.Res {
	list := s.mld.GetMusicList(ml.Mid)
	songs := make([]model.Song,0)
	for _, v := range list {
		song := s.songDao.GetSongInfo(v.Iid)
		songs = append(songs, *song)
	}
	return tool.GetGoodResult(songs)
}

// AddMusicList 添加歌单
func (s *MusicListService) AddMusicList(ml *model.MusicList) tool.Res {
	b := s.mld.AddMusicList(ml)
	if b {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("failed")
	}
}

// AddSong 添加单曲
func (s *MusicListService) AddSong(lc *model.ListContent) tool.Res {
	b := s.mld.AddSongToList(lc)
	if b {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("failed")
	}
}

// 删除单曲
func (s *MusicListService) DelSong(lc *model.ListContent) tool.Res {
	b := s.mld.DelSongFromList(lc)
	if b {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("failed")
	}
}
