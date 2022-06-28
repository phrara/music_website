package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBMgr db总会话
var DBMgr *gorm.DB

// 初始化
func init() {
	db, err := gorm.Open("mysql", "root:2001823@tcp(127.0.0.1:3306)/music_website?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("%v", errors.New("open mysql failed"))
		return
	}
	original := db.DB()
	original.SetMaxOpenConns(15)
	DBMgr = db

}
