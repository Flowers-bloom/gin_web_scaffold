package custom

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 自定义时间类型，转换从数据库查询到的时间格式
type DateTime time.Time

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