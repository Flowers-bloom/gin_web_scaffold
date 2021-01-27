package sys

import (
	"github.com/flowers-bloom/gin_web_scaffold/util/convert"
	"runtime"
)

// main.main:12
func GetCurFuncAndLine(callDepth int) string {
	pc, _, line, ok := runtime.Caller(callDepth)
	if ok {
		return runtime.FuncForPC(pc).Name() + ":" + convert.IntToString(line)
	}
	return ""
}