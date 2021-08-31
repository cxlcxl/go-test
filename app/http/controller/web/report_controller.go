package web

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Report struct{}

func (r *Report) GetReport(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	days := c.Query("days")
	appKey := c.Query("app_key")
	adsId := c.Query("ads_id")
	limitStart := ((page) - 1) * limit
	whereParams := map[string]interface{}{
		"app_key": appKey,
		"ads_id":  adsId,
		"days":    days,
	}
	where := (&tool.WhereQuery{Filter: true}).GenerateWhere(whereParams)
	counts, showList := model.ReportThirdDataDb().Show(where, limitStart, limit)
	response.Success(c, "请求成功", gin.H{
		"maps": map[string]string{
			"views":   "展示",
			"clicks":  "点击",
			"revenue": "收益",
		},
		"reports": showList,
		"total":   counts,
	})
}
