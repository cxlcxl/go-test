package logic

import (
	"encoding/json"
	"time"
)

// UnmarshalString JSON 字符串解析为字符串切片
func UnmarshalString(data string) (rs []string) {
	_ = json.Unmarshal([]byte(data), &rs)
	return
}

// TimeAgo 计算时间
func TimeAgo(t time.Time, s, format string) string {
	if format == "" {
		format = "2006-01-02"
	}
	d, _ := time.ParseDuration(s)
	date := t.Add(d).Format(format)
	return date
}
