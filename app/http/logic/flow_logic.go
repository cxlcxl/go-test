package logic

import (
	"encoding/json"
	"goskeleton/app/global/variable"
	"goskeleton/app/model"
	modelApi "goskeleton/app/model/api"
	"goskeleton/app/model/tool"
	"log"
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
		flows[i].ConfTypeName = tool.ConfTypes[flow.ConfType]
		userIds[i] = flow.OperatorId
	}
	params := map[string]interface{}{"app_key": appKey}
	where := (&tool.WhereQuery{Filter: true}).GenerateWhere(params)
	// 基本信息版本
	adsRelConf := modelApi.AdsRelConfDB().GetsBy(where)
	// 操作人
	users := model.CreateUserFactory().GetsByIds(userIds)
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

// UnmarshalFlowConf 解析配置的JSON格式数据
func UnmarshalFlowConf(detail *modelApi.FlowConfModel) (flowDetail map[string]interface{}) {
	flowDetail = map[string]interface{}{
		"flow_id":           detail.Id,
		"app_key":           detail.AppKey,
		"conf_type":         detail.ConfType,
		"conf_name":         detail.ConfName,
		"user_conf_type":    detail.UserConfType,
		"user_conf":         detail.UserConf,
		"channel_conf_type": detail.ChannelConfType,
		"channel_conf":      detail.ChannelConf,
		"area_conf_type":    detail.AreaConfType,
		"area_conf":         detail.AreaConf,
		"game_conf_type":    detail.GameConfType,
		"game_conf":         detail.GameConf,
		"sdk_conf_type":     detail.SdkConfType,
		"sdk_conf":          detail.SdkConf,
		"brand_conf_type":   detail.BrandConfType,
		"brand_conf":        detail.BrandConf,
		"sys_conf_type":     detail.SysConfType,
		"sys_conf":          detail.SysConf,
		"idfa_conf_type":    detail.IdfaConfType,
		"idfa_conf":         detail.IdfaConf,
		"rel_version":       detail.RelVersion,
	}
	if v, _ := flowDetail["area_conf_type"]; v == 1 {
		flowDetail["area_conf"] = combineFlowArea(detail.AreaConf)
	}
	if v, _ := flowDetail["channel_conf_type"]; v == 1 {
		flowDetail["channel_conf"] = combineFlowChannel(detail.ChannelConf)
	}
	if v, _ := flowDetail["brand_conf_type"]; v == 1 {
		brand := UnmarshalString(detail.BrandConf)
		brandConf := make([]map[string]interface{}, len(brand))
		for i, s := range brand {
			brandConf[i] = map[string]interface{}{"id": s, "name": s, "level": 1, "parent_id": 0}
		}
		flowDetail["brand_conf"] = brandConf
	}
	if v, _ := flowDetail["user_conf_type"]; v == 1 {
		flowDetail["user_conf"] = UnmarshalString(detail.UserConf)
	}
	if v, _ := flowDetail["game_conf_type"]; v == 1 {
		flowDetail["game_conf"] = UnmarshalString(detail.GameConf)
	}
	if v, _ := flowDetail["sdk_conf_type"]; v == 1 {
		flowDetail["sdk_conf"] = UnmarshalString(detail.SdkConf)
	}
	if v, _ := flowDetail["sys_conf_type"]; v == 1 {
		var sysConf modelApi.SysConf
		_ = json.Unmarshal([]byte(detail.SysConf), &sysConf)
		flowDetail["sys_conf"] = sysConf
	}
	adInfo := combineAdInfo(detail.Id, detail.AppKey, detail.RelVersion)
	for t, name := range tool.MAdSubType {
		if _, ok := adInfo[t]; !ok {
			adInfo[t] = map[string]interface{}{"ad_type": t, "name": name, "status": 0}
		}
	}
	flowDetail["ad_info"] = adInfo
	return
}

// 流量配置组合广告配置信息
func combineAdInfo(flowId int64, appKey string, relVersion int) map[int]map[string]interface{} {
	where := (&tool.WhereQuery{}).GenerateWhere(map[string]interface{}{"flow_id": flowId})
	flowAdTypeRel := modelApi.FlowAdTypeRelDB().GetsBy(where, "ad_type asc")
	if len(flowAdTypeRel) > 0 {
		adInfo := make(map[int]map[string]interface{}, 0)
		for _, relModel := range flowAdTypeRel {
			// 广告类型异常
			adTypeName, ok := tool.MAdSubType[relModel.AdType]
			if !ok {
				continue
			}
			dspAdsList, integrationAdsList := combineAdsIdsList(relModel.AppKey, relModel.AdType, relVersion)
			adInfo[relModel.AdType] = map[string]interface{}{
				"status":               relModel.Status,
				"ad_type":              relModel.AdType,
				"name":                 adTypeName,
				"dsp_ads_list":         dspAdsList,
				"integration_ads_list": integrationAdsList,
			}
			// 填充广告类型数据
			fillAdType(adInfo, relModel)
			// 因逻辑相同可同时填充一般广告商、优先广告商、DSP广告商
			fillGeneratePriorityDsp(adInfo, relModel, flowId)
			if relModel.IsAppRel == 1 {
				fillAppRel(adInfo, relModel, flowId, appKey)
			}
			if relModel.IsBlockPolicy == 1 {
				fillPosPolicy(adInfo, relModel, flowId, appKey)
			}
		}

		return adInfo
	}
	return nil
}

// 填充数据
func fillAdType(adInfo map[int]map[string]interface{}, flowAdTypeRel *modelApi.FlowAdTypeRelModel) {
	adInfo[flowAdTypeRel.AdType]["is_priority"] = flowAdTypeRel.IsPriority
	adInfo[flowAdTypeRel.AdType]["is_delay"] = flowAdTypeRel.IsDelay
	adInfo[flowAdTypeRel.AdType]["time"] = flowAdTypeRel.Time
	adInfo[flowAdTypeRel.AdType]["is_use_dsp"] = flowAdTypeRel.IsUseDsp
	adInfo[flowAdTypeRel.AdType]["price"] = flowAdTypeRel.Price
	adInfo[flowAdTypeRel.AdType]["is_app_rel"] = flowAdTypeRel.IsAppRel
	adInfo[flowAdTypeRel.AdType]["is_block_policy"] = flowAdTypeRel.IsBlockPolicy
	adInfo[flowAdTypeRel.AdType]["is_default"] = flowAdTypeRel.IsDefault
}
func fillGeneratePriorityDsp(adInfo map[int]map[string]interface{}, flowAdTypeRel *modelApi.FlowAdTypeRelModel, flowId int64) {
	params := map[string]interface{}{"ad_type": flowAdTypeRel.AdType, "flow_id": flowId, "conf_type": []interface{}{"in", []int{1, 2, 3}}}
	where := (&tool.WhereQuery{}).GenerateWhere(params)
	flowAdsRel := modelApi.FlowAdsRelDB().GetsBy(where, "position asc")
	if len(flowAdsRel) == 0 {
		adInfo[flowAdTypeRel.AdType]["priority_list"] = []interface{}{}
		adInfo[flowAdTypeRel.AdType]["general_list"] = []interface{}{}
		adInfo[flowAdTypeRel.AdType]["dsp_list"] = []interface{}{}
	} else {
		priority := make([]map[string]interface{}, 0)
		generate := make([]map[string]interface{}, 0)
		dsp := make([]map[string]interface{}, 0)
		adsId := make([]string, 0)
		for _, relModel := range flowAdsRel {
			adsId = append(adsId, relModel.AdsId)
			// 有优先广告商且当前数据是优先广告商
			if flowAdTypeRel.IsPriority == 1 && relModel.ConfType == 2 {
				priority = append(priority, map[string]interface{}{
					"ads_id":        relModel.AdsId,
					"name":          relModel.AdsId,
					"limit_num":     relModel.LimitNum,
					"exposure_num":  relModel.ExposureNum,
					"req_limit_num": relModel.ReqLimitNum,
					"limit_time":    relModel.LimitTime,
					"position":      relModel.Position,
				})
			}
			// 当前数据是一般广告商
			if relModel.ConfType == 1 {
				generate = append(generate, map[string]interface{}{
					"ads_id":        relModel.AdsId,
					"name":          relModel.AdsId,
					"limit_num":     relModel.LimitNum,
					"exposure_num":  relModel.ExposureNum,
					"req_limit_num": relModel.ReqLimitNum,
					"limit_time":    relModel.LimitTime,
					"position":      relModel.Position,
					"weight":        relModel.Weight,
				})
			}
			// 当前数据是DSP广告商
			if relModel.ConfType == 3 {
				dsp = append(dsp, map[string]interface{}{
					"ads_id":   relModel.AdsId,
					"name":     relModel.AdsId,
					"position": relModel.Position,
				})
			}
		}
		params = map[string]interface{}{"ad_type": []interface{}{"in", []int{1, 3}}, "ads_id": []interface{}{"in", adsId}}
		where = (&tool.WhereQuery{}).GenerateWhere(params)
		adsList := modelApi.AdsListDB().GetsBy(where, "ads_id asc")
		if len(adsList) > 0 {
			list := make(map[string]*modelApi.AdsListModel, 0)
			for _, listModel := range adsList {
				list[listModel.AdsId] = listModel
			}
			// name 字段覆盖填充
			for i, relModel := range generate {
				generate[i]["name"] = list[relModel["ads_id"].(string)].Name
			}
			for i, relModel := range dsp {
				dsp[i]["name"] = list[relModel["ads_id"].(string)].Name
			}
			if flowAdTypeRel.IsPriority == 1 {
				for i, relModel := range priority {
					priority[i]["name"] = list[relModel["ads_id"].(string)].Name
				}
			}
		}
		adInfo[flowAdTypeRel.AdType]["general_list"] = generate
		adInfo[flowAdTypeRel.AdType]["dsp_list"] = dsp
		if flowAdTypeRel.IsPriority == 1 {
			adInfo[flowAdTypeRel.AdType]["priority_list"] = priority
		}
	}
}
func fillAppRel(adInfo map[int]map[string]interface{}, flowAdTypeRel *modelApi.FlowAdTypeRelModel, flowId int64, appKey string) {
	adDevPos := getAdDevPos(appKey, flowAdTypeRel.AdType)
	log.Println(adDevPos)
}
func fillPosPolicy(adInfo map[int]map[string]interface{}, flowAdTypeRel *modelApi.FlowAdTypeRelModel, flowId int64, appKey string) {

}
func getAdDevPos(appKey string, adType int) map[string]map[string]interface{} {
	where := (&tool.WhereQuery{}).GenerateWhere(map[string]interface{}{"app_key": appKey})
	app := modelApi.AdAppDB().GetBy(where)
	params := map[string]interface{}{
		"pos_key_type": tool.MAdPosType[adType],
		"app_id":       app.AppId,
		"del":          1,
	}
	where = (&tool.WhereQuery{}).GenerateWhere(params)
	adDevPos := modelApi.AdDeverPosDB().GetsBy(where, "")
	if len(adDevPos) > 0 {
		rs := make(map[string]map[string]interface{}, 0)
		for _, po := range adDevPos {
			rs[po.DeverPosKey] = map[string]interface{}{
				"pos_key":              po.DeverPosKey,
				"pos_name":             po.DeverPosName,
				"state":                po.State,
				"limit_num":            po.LimitNum,
				"rate":                 po.Rate,
				"third_party_block_id": "",
			}
		}

		return rs
	}
	return nil
}

// 组合广告商ID列表
func combineAdsIdsList(appKey string, adType, relVersion int) (dspAdsList, integrationAdsList map[string]string) {
	where := (&tool.WhereQuery{}).GenerateWhere(map[string]interface{}{"ad_sub_type": adType, "app_key": appKey, "version": relVersion})
	adsAppRel := modelApi.AdsAppRelDB().GetsBy(where, "ads_id asc")
	if len(adsAppRel) == 0 {
		return
	}
	adsId := make([]string, len(adsAppRel))
	for i, relModel := range adsAppRel {
		adsId[i] = relModel.AdsId
	}
	params := map[string]interface{}{
		"ad_type": []interface{}{"in", []int{1, 3}},
		"ads_id":  []interface{}{"in", adsId},
	}
	where = (&tool.WhereQuery{}).GenerateWhere(params)
	adsList := modelApi.AdsListDB().GetsBy(where, "ads_id asc")
	if len(adsList) == 0 {
		return
	}
	dspAdsList, integrationAdsList = make(map[string]string, 0), make(map[string]string, 0)
	for _, listModel := range adsList {
		if listModel.AdType == 3 {
			dspAdsList[listModel.AdsId] = listModel.Name
		} else {
			integrationAdsList[listModel.AdsId] = listModel.Name
		}
	}
	return
}

// 流量配置组合区域信息
func combineFlowArea(area string) []variable.FlowArea {
	areaConf := UnmarshalString(area)
	if len(areaConf) == 0 {
		return nil
	}

	areas := make([]variable.FlowArea, len(areaConf))
	flowArea := variable.GetFlowCombine()
	for i, id := range areaConf {
		for key, area := range flowArea {
			if id == key {
				areas[i] = area
				break
			}
		}
	}
	return areas
}

// 流量配置组合渠道信息
func combineFlowChannel(channel string) []map[string]interface{} {
	channelConf := UnmarshalString(channel)
	if len(channelConf) == 0 {
		return nil
	}
	channels := make([]map[string]interface{}, len(channelConf))
	where := (&tool.WhereQuery{Filter: true}).GenerateWhere(map[string]interface{}{"channel_id": []interface{}{"in", channelConf}})
	channelInfo := modelApi.ChannelDB().GetChannel(where)
	if len(channelInfo) > 0 {
		for i, info := range channelInfo {
			channels[i] = map[string]interface{}{"id": info.ChannelId, "name": info.ChannelName, "level": 2, "parent_id": info.GroupId}
		}
	}

	return channels
}
