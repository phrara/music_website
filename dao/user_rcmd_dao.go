package dao

import "MusicWebsite/model"

type UserRCMDDao struct {
}

func (u *UserRCMDDao) GetRcmdUsers(uid string) *model.UserRCMD {
	urcmd := model.NewUserRCMD(uid)
	DBMgr.Where("uid = ?", uid).First(urcmd)
	return urcmd
}