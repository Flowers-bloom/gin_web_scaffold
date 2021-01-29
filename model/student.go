package model

import (
	"github.com/flowers-bloom/gin_web_scaffold/common/custom"
	"gorm.io/gorm"
)

type Student struct {
	Id uint `gorm:"primary_key; auto_increment" json:"id"`
	Name string `gorm:"not null; type:string; size:10" json:"name" form:"name" binding:"required,max=10"`
	Sex string `gorm:"not null; type:string; size:1" json:"sex" form:"sex" binding:"required,len=1"`
	Age byte `gorm:"not null; type:tinyint" json:"age" form:"age" binding:"required,max=127"`
	CreatedAt custom.DateTime `json:"created_at"`
	UpdatedAt custom.DateTime `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}