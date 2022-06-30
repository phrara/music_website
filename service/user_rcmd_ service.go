package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
)

type UserRCMDService struct {
	urd *dao.UserRCMDDao
	ud *dao.UserDao
}

func NewUserRCMDService() *UserRCMDService {
	urs := new(UserRCMDService)
	urs.ud = new(dao.UserDao)
	urs.urd = new(dao.UserRCMDDao)
	return urs
}

// 获取推荐用户
func (urs *UserRCMDService) GetRcmdUsers(user *model.UserInfo) tool.Res {
	ur := urs.urd.GetRcmdUsers(user.Uid)
	ulist := make([]model.User,0)
	for _, v := range ur.GetTopUsers() {
		u := urs.ud.GetUserInfo(v)
		u.Password = "*********"
		ulist = append(ulist, *u)
	}
	return tool.GetGoodResult(ulist)
}
