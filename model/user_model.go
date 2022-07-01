package model

import (
	"strconv"
	"strings"

	_ "github.com/jinzhu/gorm"
)

type User struct {
	Uid          string `gorm:"primary_key;column:uid" json:"uid"`
	Username     string `gorm:"column:name" json:"name"`
	Password     string `gorm:"column:password" json:"password"`
	Gender       string `gorm:"column:gender" json:"gender"`
	Age          string `gorm:"column:age" json:"age"`
	Area         string `gorm:"column:area" json:"area"`
	RegisterTime string `gorm:"column:registerTime" json:"registerTime"`
	// "0,1,2,3,4,5,6,7,8,9" 表示喜欢歌曲的类型
	Des          string `gorm:"column:des" json:"des"`
}

// TableName 设置表名
func (u User) TableName() string {
	return "user"
}

func (u User) TransformDes() string {
	list := strings.Split(u.Des, ",")
	des := ""
	for _, v := range list {
		i, _ := strconv.ParseInt(v, 10, 64)
		switch i {
		case ROCK:
			des += "摇滚,"
		case RB:
			des += "蓝调,"
		case JAZZ:
			des += "爵士,"
		case CLASSIC:
			des += "古典,"
		case POP:
			des += "流行,"
		case RAP:
			des += "嘻哈"
		case FOLK:
			des += "民歌"
		case PURE:
			des += "纯音乐"
		case ELECTRONIC:
			des += "电子乐"
		case COUNTRYSIDE:
			des += "乡村"
		default:
			des += ""
		}
	}
	return des + "..."
}

func NewUser(uid, username, password string) *User {
	return &User{
		Uid:      uid,
		Username: username,
		Password: password,
	}
}
