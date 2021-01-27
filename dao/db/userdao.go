package db

import (
	"github.com/flowers-bloom/gin_web_scaffold/common/db"
	"github.com/flowers-bloom/gin_web_scaffold/common/logger"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"time"
)

func GetUserById(id int) model.User {
	var user model.User

	result := db.GetDB().First(&user, id)
	if result.Error != nil {
		logger.GetLogger().ErrorL(result.Error)
	}

	return user
}

func Add(user model.User) {
	cur := model.DateTime(time.Now())
	user.CreatedAt = cur
	user.UpdatedAt = cur

	result := db.GetDB().Create(&user)
	if result.Error != nil {
		logger.GetLogger().ErrorL(result.Error)
	}
}

func Delete(user model.User) {
	db.GetDB().Delete(&user)
}

func Update(user model.User) {
	db.GetDB().Save(&user)
}