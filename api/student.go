package api

import (
	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/common/logger"
	"github.com/flowers-bloom/gin_web_scaffold/dao"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"github.com/flowers-bloom/gin_web_scaffold/util/convert"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentApi struct {
}

var studentDao dao.StudentDao

func (s *StudentApi) GetStudentById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	c.JSON(200, common.Success(0, "查询成功", studentDao.GetStudentById(id)))
}

func (s *StudentApi) GetAll(c *gin.Context) {
	c.JSON(200, common.Success(0, "查询成功", studentDao.GetAll()))
}

func (s *StudentApi) Register(c *gin.Context) {
	var Student model.Student
	if err := c.ShouldBind(&Student); err != nil {
		logger.GetLogger().ErrorL(err)
		c.JSON(http.StatusBadRequest, common.Error(http.StatusBadRequest, "参数不合法", err))
		return
	}

	studentDao.Add(Student)
	c.JSON(200, common.Success(0, "注册成功", nil))
}

func (s *StudentApi) DeleteById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	studentDao.DeleteById(id)

	c.JSON(200, common.Success(0, "删除成功", nil))
}

func (s *StudentApi) Update(c *gin.Context) {
	var Student model.Student
	if err := c.ShouldBind(&Student); err != nil {
		panic(err)
	}
	studentDao.Update(Student)

	c.JSON(200, common.Success(0, "更新成功", Student))
}