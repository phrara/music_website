package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
)

type MusicListService struct {
	mld *dao.MusicListDao
	songDao *dao.SongDao
	sgrDao *dao.SingerDao
}

func NewMusicListService() *MusicListService {
	mls := new(MusicListService)
	mls.mld = new(dao.MusicListDao)
	mls.songDao = new(dao.SongDao)
	mls.sgrDao = new(dao.SingerDao)
	return mls
}

// 获取某用户总歌单
func (s *MusicListService) GetOnesMusicList(userInfo *model.UserInfo) tool.Res {
	list := s.mld.GetMusicListsByUid(userInfo.Uid)
	return tool.GetGoodResult(list)
}

// 获取总歌单
func (m *MusicListService) GetMusicLists() tool.Res {
	return tool.GetGoodResult(m.mld.GetMusicLists(100)) 
}

// GetMusicList 获取歌单
func (s *MusicListService) GetMusicList(ml *model.MusicList) tool.Res {
	list := s.mld.GetMusicList(ml.Mid)
	infoList := make([]model.SongInfo,0)
	for _, v := range list {
		song := s.songDao.GetSongInfo(v.Iid)
		sn := s.sgrDao.GetSingerInfo(song.SingerId).SingerName
		si := model.NewSongInfo(song, sn)
		infoList = append(infoList, *si)
	}
	return tool.GetGoodResult(infoList)
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
