package api_ctl

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller"
	"goskeleton/app/http/logic"
	modelApi "goskeleton/app/model/api"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/response"
)

type AdApp struct{}

func (a *AdApp) GetAppList(c *gin.Context) {
	values := controller.GetQueries(c, []string{"app_name", "platform"})
	limitStart, pageSize := controller.GetPage(c)
	whereParams := map[string]interface{}{
		"is_check": tool.UserCheckPass,
		"platform": values["platform"].(string),
		"or": map[string]interface{}{
			"app_name": []string{"like", values["app_name"].(string)},
			"app_key":  []string{"like", values["app_name"].(string)},
		},
	}
	where := (&tool.WhereQuery{Filter: true}).GenerateWhere(whereParams)
	showList, counts := modelApi.AdAppDB().AppList(where, limitStart, pageSize)
	if counts > 0 {
		// 组合是否配置参数 is_config
		showList = logic.CombineAppConfig(showList)
	}
	response.Success(c, "请求成功", gin.H{
		"list":  showList,
		"total": counts,
	})
}

func (a *AdApp) GetFlowList(c *gin.Context) {
	appKey := c.Query("app_key")
	showList := modelApi.FlowConfDB().GetFlowByAppKey(appKey)
	if len(showList) > 0 {
		showList = logic.CombineFlowConfig(appKey, showList)
	}
	response.Success(c, "请求成功", showList)
}

func (a *AdApp) DeleteFlowConfig(c *gin.Context) {
	flowId := c.GetFloat64("id")
	if ok := modelApi.FlowConfDB().DeleteFlowById(int64(flowId)); !ok {
		response.Fail(c, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, "")
		return
	}

	response.Success(c, "删除成功", "")
}

// FlowConfigDetail 配置详情
func (a *AdApp) FlowConfigDetail(c *gin.Context) {
	values := controller.GetValues(c, []string{"flow_id.float64", "app_key", "conf_type.float64"})
	flowConf := modelApi.FlowConfDB().GetFlowConfById(int64(values["flow_id"].(float64)))
	response.Success(c, consts.CurdStatusOkMsg, logic.UnmarshalFlowConf(flowConf))
}
