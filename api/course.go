package api

import (
	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/dao"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"github.com/flowers-bloom/gin_web_scaffold/util/convert"
	"github.com/gin-gonic/gin"
)

type CourseAPI struct {
}

var courseDao dao.CourseDao

func (ca *CourseAPI) GetCourseById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	c.JSON(200, common.Success(0, "查询成功", courseDao.GetCourseById(id)))
}

func (ca *CourseAPI) GetAll(c *gin.Context) {
	c.JSON(200, common.Success(0, "查询成功", courseDao.GetAll()))
}

func (ca *CourseAPI) Add(c *gin.Context) {
	var Course model.Course
	if err := c.ShouldBind(&Course); err != nil {
		panic(err)
	}

	courseDao.Add(Course)
	c.JSON(200, common.Success(0, "添加成功", nil))
}

func (ca *CourseAPI) DeleteById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	courseDao.DeleteById(id)

	c.JSON(200, common.Success(0, "删除成功", nil))
}

func (ca *CourseAPI) Update(c *gin.Context) {
	var Course model.Course
	if err := c.ShouldBind(&Course); err != nil {
		panic(err)
	}
	courseDao.Update(Course)

	c.JSON(200, common.Success(0, "更新成功", Course))
}