package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller"
	"goskeleton/app/model/data"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/response"
)

type Report struct{}

func (r *Report) GetReport(c *gin.Context) {
	values := controller.GetQueries(c, []string{"day_start", "app_key", "ads_id", "day_end"})
	whereParams := map[string]interface{}{
		"app_key": values["app_key"].(string),
		"ads_id":  values["ads_id"].(string),
		"days":    []string{"between", values["day_start"].(string), values["day_end"].(string)},
	}
	where := (&tool.WhereQuery{Filter: true}).GenerateWhere(whereParams)
	showList := data.ReportThirdDataModel().Show(where)
	response.Success(c, "请求成功", gin.H{
		"maps": map[string]string{
			"third_views":  "展示",
			"third_clicks": "点击",
			"ad_income":    "收益",
		},
		"reports": showList,
	})
}
