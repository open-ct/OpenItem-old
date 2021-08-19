package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CreateDateDir 根据当前日期来创建文件夹
func CreateDateDir(Path string) string {
	t := time.Now()
	folderDate := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
	folderPath := filepath.Join(Path, folderDate)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, os.ModePerm) //0777
		os.Chmod(folderPath, os.ModePerm)
	}
	return folderPath
}

// IsDirExists 判断所给路径文件/文件夹是否存在
func IsDirExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
