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

// 管理员登录
func (as *AdminService) AdminLogin(admin *model.Admin) tool.Res {
	a := as.ad.ValidateAdmin(admin)
	if a.Id != 0 {
		a.Name = "***********"
		return tool.GetGoodResult(*a)
	} else {
		return tool.GetBadResult("failed")
	}
}

// 数据统计
func (as *AdminService) GetStatistics() tool.Res {
	// total users
	totalUsers := as.ad.GetTotalUsers()
	// total music lists
	totalMLs := as.ad.GetTotalMLs()
	// tolal songs
	totalSongs := as.ad.GetTotalSongs()
	// Male female ratio
	males, females := as.ad.GetTotalMalesAndFemales()
	
	data := tool.Res {
		"totalUsers": totalUsers,
		"totalMLs": totalMLs,
		"totalSongs": totalSongs,
		"males": males,
		"females": females,
	}

	return tool.GetGoodResult(data)
}