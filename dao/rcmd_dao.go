package dao

import "MusicWebsite/model"

type RCMDDao struct {
}

func (r *RCMDDao) GetRcmdUsers(uid string) *model.UserRCMD {
	urcmd := model.NewUserRCMD(uid)
	DBMgr.Where("uid = ?", uid).First(urcmd)
	return urcmd
}

func (r *RCMDDao) GetRcmdSongs(uid string) *model.SongRCMD {
	srcmd := model.NewSongRCMD(uid)
	DBMgr.Where("uid = ?", uid).First(srcmd)
	return srcmd
}

func (r *RCMDDao) GetSimilarSongs(iid string) *model.SongSimilar {
	ss := model.NewSongSimilar(iid)
	DBMgr.Where("iid = ?", iid).First(ss)
	return ss
}