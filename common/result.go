package common

import (
	"github.com/gin-gonic/gin"
)

// TODO: TraceId 链路日志追踪
type Result struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	TraceId string `json:"traceId"`
	Error string `json:"error,omitempty"`
}

// Success 请求成功返回
func Success(code int, msg string, data interface{}) Result {
	return Result{
		Code: code,
		Msg: msg,
		Data: data,
	}
}

// 通用错误处理
func Error(errCode int, msg string, err error) Result {
	res := Result{
		Code: errCode,
		Msg: msg,
	}

	// 生产环境隐藏错误
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}

	return res
}