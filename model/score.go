package model

import (
	"github.com/flowers-bloom/gin_web_scaffold/common/custom"
	"gorm.io/gorm"
)

type Score struct {
	Id uint `gorm:"primary_key; auto_increment" json:"id"`
	StudentId uint `gorm:"not null; type:uint" json:"student_id"`
	CourseId uint `gorm:"not null; type:uint" json:"course_id"`
	Value byte `gorm:"not null; type:tinyint" json:"value"`
	CreatedAt custom.DateTime `json:"created_at"`
	UpdatedAt custom.DateTime `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}