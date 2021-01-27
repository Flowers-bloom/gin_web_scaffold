package router

import (
	"github.com/flowers-bloom/gin_web_scaffold/api"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) *gin.Engine {
	// TODO：中间件


	// 路由
	user := router.Group("/user")
	{
		user.GET(":id", api.GetUserById)
		user.POST("register", api.Register)
		user.PUT("update", api.Update)
		user.DELETE("delete", api.Delete)
	}

	return router
}