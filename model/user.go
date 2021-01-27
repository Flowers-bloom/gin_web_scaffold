package model

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// 自定义时间类型，转换从数据库查询到的时间格式
type DateTime time.Time

type User struct {
	Id uint `gorm:"primary_key;auto_increment" json:"id"`
	UserName string `gorm:"type:string;size:25" json:"user_name"`
	Password string `gorm:"type:string;size:35" json:"password"`
	Email string `gorm:"not null;type:string;size:20" json:"email"`
	CreatedAt DateTime `json:"created_at"`
	UpdatedAt DateTime `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	s := time.Time(t).Format("2006-01-02 15:04:05")
	formatted := fmt.Sprintf("\"%s\"", s)
	return []byte(formatted), nil
}

func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time

	if time.Time(t).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}

func (t *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}