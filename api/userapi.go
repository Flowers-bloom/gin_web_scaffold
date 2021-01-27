package api

import (
	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/common/logger"
	"github.com/flowers-bloom/gin_web_scaffold/dao/db"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"github.com/flowers-bloom/gin_web_scaffold/util/convert"
	"github.com/gin-gonic/gin"
)

// TODO: 参数校验
func GetUserById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	user := db.GetUserById(id)

	c.JSON(200, common.Response{
		Code:0,
		Msg:"查询成功",
		Data:user,
	})
}

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		logger.GetLogger().ErrorL(err)
	}

	db.Add(user)

	c.JSON(200, common.Response{
		Code:0,
		Msg:"注册成功",
	})
}

func Delete(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		logger.GetLogger().ErrorL(err)
	}
	db.Delete(user)

	c.JSON(200, common.Response{
		Code:0,
		Msg:"删除成功",
	})
}

func Update(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		logger.GetLogger().ErrorL(err)
	}
	db.Update(user)

	c.JSON(200, common.Response{
		Code:0,
		Msg:"更新成功",
		Data:user,
	})
}