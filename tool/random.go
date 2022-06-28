package tool

import (
	"crypto/md5"
	"math/rand"
	"time"
)

func GetRandomId(args string) string {
	rand.Seed(time.Now().UnixNano())
	// 位数
	figure := 1 + rand.Int63n(11)

	// 初始化
	hash := md5.Sum([]byte(args))
	
	t := make([]string,11)
	t[0] = string(rune(int(hash[13]) % 2 + 48 + rand.Intn(9)))
	t[1] = string(rune(int(hash[11]) % 2 + 48 + rand.Intn(9)))
	t[2] = string(rune(int(hash[10]) % 2 + 48 + rand.Intn(9)))
	t[3] = string(rune(int(hash[8]) % 2 + 48 + rand.Intn(9)))
	t[4] = string(rune(int(hash[5]) % 2 + 48 + rand.Intn(9)))
	t[5] = string(rune(int(hash[4]) % 2 + 48 + rand.Intn(9)))
	t[6] = string(rune(int(hash[3]) % 2 + 48 + rand.Intn(9)))
	t[7] = string(rune(int(hash[1]) % 2 + 48 + rand.Intn(9)))
	t[8] = string(rune(int(hash[14]) %2 + 48 + rand.Intn(9)))
	t[9] = string(rune(int(hash[2]) %2 + 48 + rand.Intn(9)))
	t[10] = string(rune(int(hash[9]) %2 + 48 + rand.Intn(9)))

	res := ""

	for _, v := range t {
		res += v
	}

	return res[0:figure]

}