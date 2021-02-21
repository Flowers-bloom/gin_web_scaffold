package constant

import "time"

const (
	DefaultTimeFormat = "2006/01/02 15:04:05"
	PanicAndStackPrintFormat = "[Panic] %v\n[stack] %s"
	DefaultStudentTokenExpired = 5 * time.Minute
)