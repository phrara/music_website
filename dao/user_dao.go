package dao

import (
	"MusicWebsite/model"
)

type UserDao struct {
}

// GetUserInfo 查询用户信息
func (*UserDao) GetUserInfo(uid string) *model.User {
	user := model.NewUser("", "", "")
	DBMgr.Where("uid = ?", uid).First(user)
	return user
}

// ValidateUser 登录验证
func (*UserDao) ValidateUser(user *model.User) *model.User {
	DBMgr.Where("name = ? and password = ?", user.Username, user.Password).First(user)
	return user
}

// AddUser 添加新用户
func (ud *UserDao) AddUser(user *model.User) bool {
	if ud.GetUserInfo(user.Uid).Username == "" {
		DBMgr.Create(user)
		return true
	} else {
		return false
	}
}

// 删除用户
func (ud *UserDao) DelUser(uid string) bool {
	user := model.NewUser("", "", "")
	DBMgr.Where("uid = ?", uid).Delete(user)
	if ud.GetUserInfo(uid).Username == "" {
		return true
	} else {
		return false
	}
}

// 更改密码
func (ud *UserDao) UpdatePassword(user *model.User) bool {
	res := DBMgr.Model(user).Where("uid = ?",user.Uid).Update("password",user.Password)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}