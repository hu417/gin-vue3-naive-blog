package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

/* model/time.go */
const timeFormat = "2006-01-02 15:04:05"
const timeZone = "Asia/Shanghai"

type Time time.Time

// json转换为结构体
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

// 结构体转换为json
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, _ := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// 定义时间格式
func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

// 定义时间时区
func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(timeZone)
	return time.Time(t).In(loc)
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}