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

