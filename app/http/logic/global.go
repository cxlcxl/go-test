package logic

import "encoding/json"

// UnmarshalString JSON 字符串解析为字符串切片
func UnmarshalString(data string) (rs []string) {
	_ = json.Unmarshal([]byte(data), &rs)
	return
}
