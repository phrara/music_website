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
func (us *UserService) UserRegister(user *model.User) tool.Res {
	// 注册时间
	t := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	user.RegisterTime = t
	// 口令加密
	user.Password = tool.Encrypt(user.Password)
	// 用户id
	user.Uid = tool.GetRandomId(t + user.Username)
	b := us.userdao.AddUser(user)
	if b {
		return tool.GetGoodResult(*user)
	} else {
		return tool.GetBadResult("register failed")
	}

}

// UserLogin 用户登录
func (us *UserService) UserLogin(user *model.User) tool.Res {
	user.Username = ""
	// 口令加密
	user.Password = tool.Encrypt(user.Password)
	validatedUser := us.userdao.ValidateUser(user)
	if validatedUser.Username != "" {
		validatedUser.Password = "*************"
		// 设置 token
		token, err := tool.GetToken(validatedUser.Uid)
		if err != nil {
			return tool.GetBadResult("get Token failed")
		}
		return tool.GetGoodResult(*validatedUser, token)
	} else {
		return tool.GetBadResult("log in failed")
	}
}

// 用户注销
func (us *UserService) DeleteUser(user *model.UserInfo) tool.Res {
	if b := us.userdao.DelUser(user.Uid); b {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("failed")
	}
}

// 更改密码
func (us *UserService) UpdatePassword(userInfo *model.UserInfo) tool.Res {
	u := us.userdao.GetUserInfo(userInfo.Uid)
	if u.Password == tool.Encrypt(userInfo.OldPassword) {
		userInfo.NewPassword = tool.Encrypt(userInfo.NewPassword)
		b := us.userdao.UpdatePassword(model.NewUser(userInfo.Uid, "", userInfo.NewPassword))	
		if b {
			return tool.GetGoodResult(nil)
		} else {
			return tool.GetBadResult("err")
		}
	} else {
		return tool.GetBadResult("err")
	}
}


// 更改用户信息
func (us *UserService) UpdateInfo(user *model.User) tool.Res {
	us.userdao.UpdateUserInfo(user)
	return tool.GetGoodResult(nil)
}

// 获取所有用户
func (us *UserService) GetAllUsers() tool.Res {
	u := us.userdao.GetAllUsers(1000)
	return tool.GetGoodResult(u)
}