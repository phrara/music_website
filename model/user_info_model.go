package model

type UserInfo struct {
	Uid string `json:"uid"`
	OldPassword string `json:"oldPw"`
	NewPassword string `json:"newPw"`
}


