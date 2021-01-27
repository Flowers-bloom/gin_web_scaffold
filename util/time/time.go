package time

import (
	"github.com/flowers-bloom/gin_web_scaffold/constant"
	"time"
)

func GetCurrentTime() string {
	return time.Now().Format(constant.DefaultTimeFormat)
}

func GetCurrentTimeF(format string) string {
	return time.Now().Format(format)
}