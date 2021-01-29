package dao

import (
	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/common/custom"
	"github.com/flowers-bloom/gin_web_scaffold/common/db"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"time"
)

type StudentDao struct {
}

func (s *StudentDao) GetStudentById(id int) model.Student {
	var Student model.Student

	result := db.GetDB().First(&Student, id)
	if result.Error != nil {
		panic(result.Error)
	}

	return Student
}

func (s *StudentDao) GetAll() common.PageResult {
	var students []model.Student

	result := db.GetDB().Find(&students)
	if result.Error != nil {
		panic(result.Error)
	}

	rows := make([]interface{}, len(students))
	for i, student := range students {
		rows[i] = student
	}

	return common.PageResult{
		Total:result.RowsAffected,
		Rows:rows,
	}
}

func (s *StudentDao) Add(student model.Student) {
	cur := custom.DateTime(time.Now())
	student.CreatedAt = cur
	student.UpdatedAt = cur

	result := db.GetDB().Create(&student)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (s *StudentDao) DeleteById(id int) {
	db.GetDB().Delete(&model.Student{}, id)
}

func (s *StudentDao) Update(student model.Student) {
	db.GetDB().Save(&student)
}