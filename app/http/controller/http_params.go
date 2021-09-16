package controller

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"strconv"
	"strings"
)

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

// GetQueries 批量获取参数 string/int/float64/int64 前提：有验证器处理的值
// key 传入方式 user_name.string，只传递 user_name 默认 string 类型
func GetQueries(c *gin.Context, keys []string) map[string]interface{} {
	values := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		sKey := strings.Split(key, ".")
		if len(sKey) == 1 {
			values[sKey[0]] = c.Query(sKey[0])
		} else {
			v := c.Query(sKey[0])
			switch strings.ToLower(sKey[1]) {
			case "string":
				values[sKey[0]] = v
			case "int":
				val, err := strconv.Atoi(v)
				if err != nil {
					values[sKey[0]] = 0
				} else {
					values[sKey[0]] = val
				}
			case "uint":
				val, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					values[sKey[0]] = 0
				} else {
					values[sKey[0]] = val
				}
			case "float64":
				val, err := strconv.ParseFloat(v, 64)
				if err != nil {
					values[sKey[0]] = 0.0
				} else {
					values[sKey[0]] = val
				}
			case "int64":
				val, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					values[sKey[0]] = 0
				} else {
					values[sKey[0]] = val
				}
			default:
				values[sKey[0]] = v
			}
		}
	}
	return values
}

// GetForms 批量获取参数 string/int/float64/int64 前提：有验证器处理的值
// key 传入方式 user_name.string，只传递 user_name 默认 string 类型
func GetForms(c *gin.Context, keys []string) map[string]interface{} {
	values := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		sKey := strings.Split(key, ".")
		if len(sKey) == 1 {
			values[sKey[0]] = c.PostForm(sKey[0])
		} else {
			v := c.PostForm(sKey[0])
			switch strings.ToLower(sKey[1]) {
			case "string":
				values[sKey[0]] = v
			case "int":
				val, err := strconv.Atoi(v)
				if err != nil {
					values[sKey[0]] = 0
				} else {
					values[sKey[0]] = val
				}
			case "uint":
				val, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					values[sKey[0]] = 0
				} else {
					values[sKey[0]] = val
				}
			case "float64":
				val, err := strconv.ParseFloat(v, 64)
				if err != nil {
					values[sKey[0]] = 0.0
				} else {
					values[sKey[0]] = val
				}
			case "int64":
				val, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					values[sKey[0]] = 0
				} else {
					values[sKey[0]] = val
				}
			default:
				values[sKey[0]] = v
			}
		}
	}
	return values
}

// GetPage 快速获取分页
func GetPage(c *gin.Context) (pageStart, limit int) {
	page, err := strconv.Atoi(c.Query(consts.PageKey))
	if err != nil {
		page = 1
	}
	limit, err = strconv.Atoi(c.Query(consts.PageLimit))
	if err != nil {
		limit = 15
	}
	pageStart = (page - 1) * limit
	return
}
