package data

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller"
	"goskeleton/app/http/logic"
	"goskeleton/app/model/data"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/response"
	"log"
	"time"
)

type Report struct {
}

var (
	filterFields = []string{"start_date", "end_date", "dims", "compare", "theader", "kpis", "is_custom", "app_type", "settlement_type"}
)

func (r *Report) GetReport(c *gin.Context) {
	values := controller.GetQueries(c, []string{"day_start", "app_key", "ads_id", "day_end"})
	whereParams := map[string]interface{}{
		"app_key": values["app_key"].(string),
		"ads_id":  values["ads_id"].(string),
		"days":    []string{"between", values["day_start"].(string), values["day_end"].(string)},
	}
	where := (&tool.WhereQuery{Filter: true}).GenerateWhere(whereParams)
	showList := data.ReportThirdDataDB().Show(where)
	response.Success(c, "请求成功", gin.H{
		"maps": map[string]string{
			"third_views":  "展示",
			"third_clicks": "点击",
			"ad_income":    "收益",
		},
		"reports": showList,
	})
}

// ExternalData 对外数据
func (r *Report) ExternalData(c *gin.Context) {

}

// AdReport 广告报表
func (r *Report) AdReport(c *gin.Context) {
	params := controller.GetQueries(c, []string{"start_date", "end_date"})
	if params["start_date"] == "" && params["end_date"] == "" {
		params["start_date"] = logic.TimeAgo(time.Now(), "-168h", "")
		params["end_date"] = logic.TimeAgo(time.Now(), "-24h", "")
	}
	whereParams := map[string]interface{}{
		"start_date": params["start_date"],
		"end_date":   params["end_date"],
	}
	userId, _ := c.Get("user_id")
	log.Println(whereParams, data.GetFilterFields(), userId)
}
