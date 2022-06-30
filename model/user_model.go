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
			des += "Rock,"
		case RB:
			des += "R&B,"
		case JAZZ:
			des += "Jazz,"
		case CLASSIC:
			des += "Classic,"
		case POP:
			des += "Pop,"
		case RAP:
			des += "Rap"
		case FOLK:
			des += "Folk"
		case PURE:
			des += "Pure"
		case ELECTRONIC:
			des += "Electronic"
		case COUNTRYSIDE:
			des += "Country"
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
