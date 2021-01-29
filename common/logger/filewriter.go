package logger

import (
	"errors"
	"fmt"
	"github.com/flowers-bloom/gin_web_scaffold/util/convert"
	"github.com/flowers-bloom/gin_web_scaffold/util/file"
	"github.com/flowers-bloom/gin_web_scaffold/util/time"
	"os"
	"strings"
)

var ErrFileNotCreated = errors.New("file has not created")

type FileWriter struct {
	file *os.File
	maxLength int64 // byte
	flag int
}

func (fw *FileWriter) setFile(fileName string) {
	var err error

	if fw.file != nil {
		file.Close(fw.file)
	}

	if fw.file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		panic(err)
	}
	fw.flag = 1
}

func (fw *FileWriter) Write(data interface{}) {
	if fw.file == nil {
		panic(ErrFileNotCreated)
	}

	// 日志滚动收集
	fw.Scroll()

	str := fmt.Sprintf("%v\n", data)
	if _, err := fw.file.WriteString(str); err != nil {
		panic(err)
	}
}

func (fw *FileWriter) NextName(fileName string) string {
	fw.flag++
	split := strings.Split(fileName, ".")
	size := len(split) // xxx.log or xxx.1.log

	if size == 2 {
		return split[0] + "." + convert.IntToString(fw.flag) + "." + split[1]
	}
	return split[0] + "." + convert.IntToString(fw.flag) + "." + split[2]
}

func (fw *FileWriter) Scroll() {
	size := file.GetSize(fw.file.Name())
	if size >= fw.maxLength {
		origin := fw.file.Name()
		nextName := fw.NextName(origin)
		for file.Exist(nextName) {
			nextName = fw.NextName(origin)
		}

		// 关闭原文件
		file.Close(fw.file)
		fw.file = nil

		// 重命名原文件
		file.Rename(origin, nextName)
		fmt.Println(origin, "rename to", nextName)

		// 创建新文件
		fw.setFile(time.GetCurrentTimeF("2006-01-02") + ".log")
	}
}