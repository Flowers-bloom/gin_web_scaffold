package main

import (
	"github.com/flowers-bloom/gin_web_scaffold/config"
	"github.com/flowers-bloom/gin_web_scaffold/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// 初始化本地配置文件
	config.InitLocalConf()

	// 初始化路由
	router.Init(engine)

	// 启动 gin
	engine.Run()
}