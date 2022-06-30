package dao

import "MusicWebsite/model"

type SongDao struct {
}

// 模糊查询
func (s *SongDao) SearchFor(keyWord string) []model.Song {
	list := make([]model.Song,50)
	DBMgr.Where("song_name Like ?", "%"+keyWord+"%").Find(&list)
	return list
}

// 获取单曲列表
func (*SongDao) GetSongList(num int) []model.Song {
	list := make([]model.Song, num)
	DBMgr.Find(&list)
	return list
}

// GetSongInfo 获取单曲信息
func (*SongDao) GetSongInfo(iid string) *model.Song {
	song := model.NewSong("", "")
	DBMgr.Where("iid = ?", iid).First(song)
	return song
}

// AddSong 添加单曲
func (sd *SongDao) AddSong(song *model.Song) bool {
	if sd.GetSongInfo(song.Iid).SongName == "" {
		DBMgr.Create(song)
		return true
	} else {
		return false
	}
}

// DeleteSong 删除单曲
func (sd *SongDao) DeleteSong(iid string) bool {
	song := model.NewSong("","")
	res := DBMgr.Model(song).Where("iid = ?", iid).Update("isDelete","0")
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// 播放次数修改
func (sd *SongDao) UpdatePlayCnt(iid string, dif uint) bool {
	u := sd.GetSongInfo(iid).PlayCount
	song := model.NewSong("","")
	res := DBMgr.Model(song).Where("iid = ?", iid).Update("playcnt", u + dif)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

