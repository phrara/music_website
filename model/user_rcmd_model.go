package model

import "strings"

type UserRCMD struct {
	Uid      string `gorm:"column:uid" json:"uid"`
	TopUsers string `gorm:"column:topusers" json:"topusers"`
}

func (u UserRCMD) TableName() string {
	return "topusers"
}

func NewUserRCMD(uid string) *UserRCMD {
	return &UserRCMD{
		Uid:      uid,
		TopUsers: "",
	}
}

func (ur *UserRCMD) GetTopUsers() []string {
	return strings.Split(ur.TopUsers, ",")
}
