package dao

import (
	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/common/custom"
	"github.com/flowers-bloom/gin_web_scaffold/common/db"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"time"
)

type CourseDao struct {
}

func (c *CourseDao) GetCourseById(id int) model.Course {
	var Course model.Course

	result := db.GetDB().First(&Course, id)
	if result.Error != nil {
		panic(result.Error)
	}

	return Course
}

func (c *CourseDao) GetAll() common.PageResult {
	var Courses []model.Course

	result := db.GetDB().Find(&Courses)
	if result.Error != nil {
		panic(result.Error)
	}

	rows := make([]interface{}, len(Courses))
	for i, Course := range Courses {
		rows[i] = Course
	}

	return common.PageResult{
		Total:result.RowsAffected,
		Rows:rows,
	}
}

func (c *CourseDao) Add(Course model.Course) {
	cur := custom.DateTime(time.Now())
	Course.CreatedAt = cur
	Course.UpdatedAt = cur

	result := db.GetDB().Create(&Course)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (c *CourseDao) DeleteById(id int) {
	db.GetDB().Delete(&model.Course{}, id)
}

func (c *CourseDao) Update(Course model.Course) {
	db.GetDB().Save(&Course)
}