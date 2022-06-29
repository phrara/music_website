package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"strconv"
)

// 初始盐
const (
	prefixSalt = "BUPT_Music_Website_2022"
	suffixSalt = "YiYanDingZhen_JianDingWeiJia"
)

// Encrypt 口令加密
func Encrypt(password string) string {
	plainText := ""
	plainText = prefixSalt + password + suffixSalt
	saltBytes := md5.Sum([]byte(plainText))
	// 用户专属盐
	salt := string(saltBytes[0:])

	plainText = salt + password + salt
	cipherBytes := sha256.Sum256([]byte(plainText))
	cipher := ""
	for _, cipherByte := range cipherBytes {
		cipher += strconv.FormatInt(int64(cipherByte), 16)
	}
	return cipher
}
