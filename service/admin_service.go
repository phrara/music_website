package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
)

type AdminService struct {
	ad *dao.AdminDao
}

func NewAdminService() *AdminService{
	as := new(AdminService)
	as.ad = new(dao.AdminDao)
	return as
}


func (as *AdminService) AdminLogin(admin *model.Admin) bool {
	a := as.ad.ValidateAdmin(admin)
	if a.Id != 0 {
		return true
	} else {
		return false
	}
}