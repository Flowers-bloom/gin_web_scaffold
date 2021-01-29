package api

import (
	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/dao"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"github.com/flowers-bloom/gin_web_scaffold/util/convert"
	"github.com/gin-gonic/gin"
)

type CourseApi struct {
}

var courseDao dao.CourseDao

func (ca *CourseApi) GetCourseById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	c.JSON(200, common.Success(0, "查询成功", courseDao.GetCourseById(id)))
}

func (ca *CourseApi) GetAll(c *gin.Context) {
	c.JSON(200, common.Success(0, "查询成功", courseDao.GetAll()))
}

func (ca *CourseApi) Add(c *gin.Context) {
	var Course model.Course
	if err := c.ShouldBind(&Course); err != nil {
		panic(err)
	}

	courseDao.Add(Course)
	c.JSON(200, common.Success(0, "添加成功", nil))
}

func (ca *CourseApi) DeleteById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	courseDao.DeleteById(id)

	c.JSON(200, common.Success(0, "删除成功", nil))
}

func (ca *CourseApi) Update(c *gin.Context) {
	var Course model.Course
	if err := c.ShouldBind(&Course); err != nil {
		panic(err)
	}
	courseDao.Update(Course)

	c.JSON(200, common.Success(0, "更新成功", Course))
}