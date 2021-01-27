package db

import (
	"fmt"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	if DB == nil {
		Init(GetMysqlDSN())
	}
	return DB
}

func Init(connString string) {
	var err error

	DB, err = gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	Migrate()
}

// 数据生成方法
// 修改数据库中模型，和 model 中定义的保持一致
func Migrate() {
	if DB != nil {
		DB.AutoMigrate(&model.User{})
	}
}

func GetMysqlDSN() string {
	host := viper.GetString("mysql.host")
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	dbName := viper.GetString("mysql.db_name")

	return user + ":" + pwd + "@(" + host + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
}