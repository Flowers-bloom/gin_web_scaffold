package model

import (
	"github.com/flowers-bloom/gin_web_scaffold/common/custom"
	"gorm.io/gorm"
)

type Course struct {
	Id uint `gorm:"primary_key; auto_increment" json:"id"`
	Name string `gorm:"not null; type:string; size:20" json:"name"`
	ClassHour byte `gorm:"not null; type:tinyint" json:"class_hour"`
	Teacher string `gorm:"not null; type:string; size:20" json:"teacher"`
	CreatedAt custom.DateTime `json:"created_at"`
	UpdatedAt custom.DateTime `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}