package common

import (
	"github.com/flowers-bloom/gin_web_scaffold/common/logger"
	"github.com/gin-gonic/gin"
)

// TODO: TraceId 链路日志追踪
type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data, omitempty"`
	TraceId string `json:"traceId"`
	Error string `json:"error, omitempty"`
}

// 请求成功返回
func Success(code int, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Msg: msg,
		Data: data,
	}
}

// 通用错误处理
func Error(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg: msg,
	}

	// 打印错误到控制台
	logger.GetLogger().ErrorL(err)

	// 生产环境隐藏错误
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}

	return res
}