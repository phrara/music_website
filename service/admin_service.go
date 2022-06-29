package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
)

type AdminService struct {
	ad *dao.AdminDao
}

func NewAdminService() *AdminService{
	as := new(AdminService)
	as.ad = new(dao.AdminDao)
	return as
}


func (as *AdminService) AdminLogin(admin *model.Admin) tool.Res {
	a := as.ad.ValidateAdmin(admin)
	if a.Id != 0 {
		a.Name = "***********"
		a.Password = "************"
		return tool.GetGoodResult(*a)
	} else {
		return tool.GetBadResult("failed")
	}
}