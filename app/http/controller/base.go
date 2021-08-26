package controller

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// GetStrings 批量获取参数 前提：有验证器处理的值
func GetStrings(c *gin.Context, keys []string) map[string]string {
	values := make(map[string]string, len(keys))
	for _, key := range keys {
		values[key] = c.GetString(key)
	}
	return values
}

// GetValues 批量获取参数 string/int/float64/int64 前提：有验证器处理的值
// key 传入方式 user_name.string，只传递 user_name 默认 string 类型
func GetValues(c *gin.Context, keys []string) map[string]interface{} {
	values := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		sKey := strings.Split(key, ".")
		if len(sKey) == 1 {
			values[sKey[0]] = c.GetString(sKey[0])
		} else {
			switch strings.ToLower(sKey[1]) {
			case "string":
				values[sKey[0]] = c.GetString(sKey[0])
			case "int":
				values[sKey[0]] = c.GetInt(sKey[0])
			case "uint":
				values[sKey[0]] = c.GetUint(sKey[0])
			case "float64":
				values[sKey[0]] = c.GetFloat64(sKey[0])
			case "int64":
				values[sKey[0]] = c.GetInt64(sKey[0])
			default:
				values[sKey[0]] = c.GetString(sKey[0])
			}
		}
	}
	return values
}
