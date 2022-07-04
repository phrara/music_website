package dao

import (
	"MusicWebsite/model"
)

type MusicListDao struct {
}

// GetMusicList 获取某歌单
func (m *MusicListDao) GetMusicList(mid int) []model.ListContent {
	list := make([]model.ListContent, 10)
	DBMgr.Where("mid = ?", mid).Find(&list)
	return list
}

// 获取某用户的总歌单
func (m *MusicListDao) GetMusicListsByUid(uid string) []model.MusicList {
	mls := make([]model.MusicList, 10)
	DBMgr.Where("uid = ?", uid).Find(&mls)
	return mls
}

// 获取总歌单
func (m *MusicListDao) GetMusicLists(num int) []model.MusicList {
	mls := make([]model.MusicList, num)
	DBMgr.Order("mid desc").Find(&mls)
	return mls
}


// AddSongToList 添加单曲到歌单
func (m *MusicListDao) AddSongToList(lc *model.ListContent) bool {
	b, _ := m.CheckContent(lc.Mid, lc.Iid)
	if !b {
		DBMgr.Create(lc)
		return true
	} else {
		return false
	}
}

// DelSongFromList 从歌单中删除单曲
func (m *MusicListDao) DelSongFromList(lc *model.ListContent) bool {
	b, id := m.CheckContent(lc.Mid, lc.Iid)
	if b {
		DBMgr.Where("id = ?", id).Delete(&model.ListContent{})
		return true
	} else {
		return false
	}
}

// CheckContent 查询歌单中某单曲是否存在
func (m *MusicListDao) CheckContent(mid int, iid string) (bool, int) {
	lc := model.NewListContent(0, "")
	DBMgr.Where("mid = ? and iid = ?", mid, iid).First(lc)
	if lc.Id <= 0 {
		return false, -1
	} else {
		return true, lc.Id
	}
}

// AddMusicList 创建歌单
func (m *MusicListDao) AddMusicList(ml *model.MusicList) bool {
	mlist := model.NewMusicList(0, "", "")
	DBMgr.Where("mid = ?", ml.Mid).First(mlist)
	if mlist.ListName == "" {
		DBMgr.Create(ml)
		return true
	} else {
		return false
	}
}


// 删除歌单
func (m *MusicListDao) DeleteMuisicList(mid int) bool {
	ml := model.NewMusicList(0, "", "")
	lc := model.NewListContent(0, "")
	DBMgr.Where("mid = ?", mid).Delete(lc)
	res := DBMgr.Where("mid = ?", mid).Delete(ml)
	if res.RowsAffected >= 0 {
		return true
	} else {
		return false
	}
}