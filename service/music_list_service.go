package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
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
func (s *MusicListService) GetOnesMusicList(userInfo *model.UserInfo) []model.MusicList {
	list := s.mld.GetMusicListsByUid(userInfo.Uid)
	return list
}

// 获取总歌单
func (m *MusicListService) GetMusicLists() []model.MusicList {
	return m.mld.GetMusicLists()
}

// GetMusicList 获取歌单
func (s *MusicListService) GetMusicList(ml *model.MusicList) []model.Song {
	list := s.mld.GetMusicList(ml.Mid)
	songs := make([]model.Song,0)
	for _, v := range list {
		song := s.songDao.GetSongInfo(v.Iid)
		songs = append(songs, *song)
	}
	return songs
}

// AddMusicList 添加歌单
func (s *MusicListService) AddMusicList(ml *model.MusicList) bool {
	b := s.mld.AddMusicList(ml)
	if b {
		return true
	} else {
		return false
	}
}

// AddSong 添加单曲
func (s *MusicListService) AddSong(lc *model.ListContent) bool {
	b := s.mld.AddSongToList(lc)
	if b {
		return true
	} else {
		return false
	}
}

// 删除单曲
func (s *MusicListService) DelSong(lc *model.ListContent) bool {
	b := s.mld.DelSongFromList(lc)
	if b {
		return true
	} else {
		return false
	}
}
