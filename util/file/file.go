package file

import (
	"fmt"
	"os"
)

func GetSize(fileName string) int64 {
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
	}

	return info.Size()
}

func Close(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func Rename(oldPath string, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println(err)
	}
}

func Exist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return true
}