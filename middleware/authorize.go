package middleware

import (
	"fmt"

	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/common/redis"
	"github.com/flowers-bloom/gin_web_scaffold/constant"
	"github.com/gin-gonic/gin"
)

// Authorize 对用户进行身份认证
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				fmt.Println(p)
			}
		}()

		token := c.Request.Header.Get("Token")
		expired := redis.TTL(token)
		if expired == -2 {
			c.JSON(401, common.Success(401, "请先登录", nil))
			c.Abort()
		}

		// Token 延时
		if expired <= (constant.DefaultStudentTokenExpired >> 1) {
			redis.Expire(token, constant.DefaultStudentTokenExpired)
		}

		c.Set("Token", token)
		c.Next()
	}
}
