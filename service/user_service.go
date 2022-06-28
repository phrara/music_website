package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
	"strconv"
	"time"
)

// UserService 用户功能
type UserService struct {
	userdao *dao.UserDao
}

func NewUserService() *UserService {
	us := new(UserService)
	us.userdao = new(dao.UserDao)
	return us
}

// UserRegister 用户注册
func (us *UserService) UserRegister(user *model.User) bool {
	// 注册时间
	t := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	user.RegisterTime = t
	// 口令加密
	user.Password = tool.Encrypt(user.Password)
	// 用户id
	user.Uid = tool.GetRandomId(t)
	b := us.userdao.AddUser(user)
	if b {
		return true
	} else {
		return false
	}

}

// UserLogin 用户登录
func (us *UserService) UserLogin(user *model.User) (bool, *model.User) {
	user.Uid = ""
	// 口令加密
	user.Password = tool.Encrypt(user.Password)
	validatedUser := us.userdao.ValidateUser(user)
	if validatedUser.Uid != "" {
		return true, user
	} else {
		return false, nil
	}
}

// 用户注销
func (us *UserService) DeleteUser(user *model.UserInfo) (bool) {
	if b := us.userdao.DelUser(user.Uid); b {
		return true
	} else {
		return false
	}
}

// 更改密码
func (us *UserService) UpdatePassword(userInfo *model.UserInfo) bool {
	u := us.userdao.GetUserInfo(userInfo.Uid)
	if u.Password == tool.Encrypt(userInfo.OldPassword) {
		userInfo.NewPassword = tool.Encrypt(userInfo.NewPassword)
		b := us.userdao.UpdatePassword(model.NewUser(userInfo.Uid, "", userInfo.NewPassword))	
		if b {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}


