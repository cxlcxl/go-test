package logic

import (
	"goskeleton/app/model"
	modelApi "goskeleton/app/model/api"
	"goskeleton/app/model/tool"
)

// CombineAppConfig 应用配置信息组合配置
func CombineAppConfig(apps []*modelApi.AppBaseInfo) []*modelApi.AppBaseInfo {
	appKeys := make([]string, len(apps))
	for i, app := range apps {
		appKeys[i] = app.AppKey
	}
	flowConf := modelApi.FlowConfDB().AppFlowConfig(appKeys)
	if len(flowConf) > 0 {
		tmp := make(map[string]int, len(flowConf))
		for _, set := range flowConf {
			tmp[set.AppKey] = set.Counts
		}
		for i, app := range apps {
			if _, ok := tmp[app.AppKey]; ok {
				apps[i].IsConfig = 1
			}
		}
	}

	return apps
}

// CombineFlowConfig 应用配置信息组合配置
func CombineFlowConfig(appKey string, flows []*modelApi.FlowList) []*modelApi.FlowList {
	userIds := make([]int64, len(flows))
	for i, flow := range flows {
		flows[i].ConfTypeName = model.ConfTypes[flow.ConfType]
		userIds[i] = flow.OperatorId
	}
	params := map[string]interface{}{"app_key": appKey}
	where := (&tool.WhereQuery{Filter: true}).GenerateWhere(params)
	// 基本信息版本
	adsRelConf := modelApi.AdsRelConfDB().GetsBy(where)
	// 操作人
	users := model.CreateUserFactory("").GetsByIds(userIds)
	confTmp := make(map[int]string, len(adsRelConf))
	userTmp := make(map[int64]string, len(users))
	if len(adsRelConf) > 0 {
		for _, set := range adsRelConf {
			confTmp[set.Version] = set.ConfName
		}
	}
	if len(users) > 0 {
		for _, set := range users {
			userTmp[set.Id] = set.UserName
		}
	}
	for i, flow := range flows {
		if name, ok := confTmp[flow.RelVersion]; ok {
			flows[i].RelVersionName = name
		}
		if name, ok := userTmp[flow.OperatorId]; ok {
			flows[i].Operator = name
		}
	}
	return flows
}
