package api

import (
	"net/http"

	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/common/logger"
	"github.com/flowers-bloom/gin_web_scaffold/common/redis"
	"github.com/flowers-bloom/gin_web_scaffold/constant"
	"github.com/flowers-bloom/gin_web_scaffold/dao"
	"github.com/flowers-bloom/gin_web_scaffold/dto"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"github.com/flowers-bloom/gin_web_scaffold/util/convert"
	"github.com/flowers-bloom/gin_web_scaffold/util/random"
	"github.com/gin-gonic/gin"
)

type StudentAPI struct {
}

var studentDao dao.StudentDao

func (s *StudentAPI) GetStudentById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	c.JSON(200, common.Success(0, "查询成功", studentDao.GetStudentById(id)))
}

func (s *StudentAPI) GetAll(c *gin.Context) {
	c.JSON(200, common.Success(0, "查询成功", studentDao.GetAll()))
}

func (s *StudentAPI) Register(c *gin.Context) {
	var Student model.Student
	if err := c.ShouldBind(&Student); err != nil {
		logger.GetLogger().ErrorL(err)
		c.JSON(http.StatusBadRequest, common.Error(http.StatusBadRequest, "参数不合法", err))
		c.Abort()
		return
	}

	studentDao.Add(Student)
	c.JSON(200, common.Success(0, "注册成功", nil))
}

func (s *StudentAPI) DeleteById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	studentDao.DeleteById(id)

	c.JSON(200, common.Success(0, "删除成功", nil))
}

func (s *StudentAPI) Update(c *gin.Context) {
	var Student model.Student
	if err := c.ShouldBind(&Student); err != nil {
		panic(err)
	}
	studentDao.Update(Student)

	c.JSON(200, common.Success(0, "更新成功", Student))
}

func (s *StudentAPI) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	if err := c.ShouldBind(&loginDto); err != nil {
		logger.GetLogger().ErrorL(err)
		c.JSON(http.StatusBadRequest, common.Error(http.StatusBadRequest, "参数不合法", err))
		c.Abort()
		return
	}

	var student model.Student
	studentDao.Exist(loginDto.Email, &student)
	if (student.Id == 0) {
		c.JSON(200, common.Success(0, "邮箱未注册", nil))
		c.Abort()
		return
	}else if (student.Pwd != loginDto.Pwd) {
		c.JSON(200, common.Success(0, "密码错误", nil))
		c.Abort()
		return
	}

	token := random.RandString(32)
	redis.SetEX(token, convert.IntToString(int(student.Id)), constant.DefaultStudentTokenExpired)
	c.JSON(200, common.Success(0, "登录成功", dto.LoginResult{
		Token: token,
		Student: student,
	}))
}