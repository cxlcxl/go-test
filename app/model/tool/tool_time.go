package tool

import (
	"database/sql/driver"
	"goskeleton/app/global/variable"
	"log"
	"time"
)

// LocalTime 时间格式化格式：2005-01-02 15:04:05
type LocalTime struct {
	time.Time
}

// LocalDate 时间格式化格式：2005-01-02
type LocalDate struct {
	time.Time // 2005-01-02
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	d := t.Format(variable.DateFormat)
	return []byte(`"` + d + `"`), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if !ok {
		switch v.(type) {
		case int64:
			value = time.Unix(v.(int64), 0)
		default:
			log.Println("时间戳格式不能解析", v)
			return nil
		}
	}
	*t = LocalTime{Time: value}
	return nil
}

func (l LocalDate) MarshalJSON() ([]byte, error) {
	d := l.Format("2006-01-02")
	return []byte(`"` + d + `"`), nil
}

func (l LocalDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if l.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return l.Time, nil
}

func (l *LocalDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if !ok {
		switch v.(type) {
		case int64:
			value = time.Unix(v.(int64), 0)
		default:
			log.Println("时间戳格式不能解析", v)
			return nil
		}
	}
	*l = LocalDate{Time: value}
	return nil
}
