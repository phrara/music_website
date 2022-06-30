package tool

import (
	"bufio"
	"fmt"
	"os"
)

// 读取文件操作
func ReadFile(filePath string) (bool, string) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, string(b)
}

// 写文件
func WriteFile(filePath string, data string) bool {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	w := bufio.NewWriter(f)
	w.WriteString(data + "\n")
	w.Flush()
	return true
}