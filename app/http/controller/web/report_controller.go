package web

import (
	"goskeleton/app/utils/response"

	"github.com/gin-gonic/gin"
)

type Report struct{}

func (r *Report) GetReport(c *gin.Context) {
	response.Success(c, "请求成功", gin.H{
		"maps": map[string]string{
			"views":   "展示",
			"clicks":  "点击",
			"revenue": "收益",
		},
		"reports": []map[string]interface{}{
			{"days": "2021-08-13", "views": 153, "revenue": 2, "clicks": 121},
			{"days": "2021-08-14", "views": 166, "revenue": 3, "clicks": 100},
			{"days": "2021-08-15", "views": 127, "revenue": 4.6, "clicks": 90},
			{"days": "2021-08-16", "views": 110, "revenue": 3.2, "clicks": 101},
			{"days": "2021-08-17", "views": 100, "revenue": 2, "clicks": 80},
			{"days": "2021-08-18", "views": 120, "revenue": 1.9, "clicks": 82},
			{"days": "2021-08-19", "views": 161, "revenue": 5, "clicks": 91},
			{"days": "2021-08-20", "views": 134, "revenue": 4.3, "clicks": 154},
			{"days": "2021-08-21", "views": 175, "revenue": 3.8, "clicks": 162},
			{"days": "2021-08-22", "views": 160, "revenue": 4.5, "clicks": 140},
			{"days": "2021-08-23", "views": 165, "revenue": 6, "clicks": 145},
		},
	})
}
