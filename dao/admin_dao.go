package dao

import (
	"MusicWebsite/model"
)

type AdminDao struct {
}

// @title	ValidateAdmin
// @description   验证管理员
// @auth      彭竑睿           时间（2022/6/28）
// @param     admin      *model.Admin      "管理员类"
// @return    admin      *model.Admin      "管理员类"
func (a *AdminDao) ValidateAdmin(admin *model.Admin) *model.Admin {
	DBMgr.Where("name = ? and password = ?",admin.Name, admin.Password).First(admin)
	return admin
}

func (a *AdminDao) GetTotalUsers() int {
	var c int
	DBMgr.Model(&model.User{}).Count(&c)
	return c
}

func (a *AdminDao) GetTotalSongs() int {
	var c int
	DBMgr.Model(&model.Song{}).Count(&c)
	return c
}

func (a *AdminDao) GetTotalMLs() int {
	var c int
	DBMgr.Model(&model.MusicList{}).Count(&c)
	return c
}

func (a *AdminDao) GetTotalMalesAndFemales() (int, int) {
	var(
		mc int
		fc int
	) 
	DBMgr.Model(&model.User{}).Where("gender = '男'").Count(&mc)
	DBMgr.Model(&model.User{}).Where("gender = '女'").Count(&fc)
	return mc, fc
}
