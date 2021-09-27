package model_api

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

func FlowAdsRelDB() *FlowAdsRelModel {
	return &FlowAdsRelModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

type FlowAdsRelModel struct {
	model.BaseModel `json:"-"`
	FlowId          int     `gorm:"column:flow_id" json:"flow_id"` // 流量配置id
	AppKey          string  `gorm:"column:app_key" json:"app_key"`
	AdType          int     `gorm:"column:ad_type" json:"ad_type"`             // 广告类型 1视频2插页
	ConfType        int     `gorm:"column:conf_type" json:"conf_type"`         // 1一般广告商2优先广告商3DSP配置广告商
	AdsId           string  `gorm:"column:ads_id" json:"ads_id"`               // 广告商id
	Position        int     `json:"position"`                                  // 位置
	LimitNum        int     `gorm:"column:limit_num" json:"limit_num"`         // 限制次数
	ExposureNum     int     `gorm:"column:exposure_num" json:"exposure_num"`   // 曝光量限制
	ReqLimitNum     int     `gorm:"column:req_limit_num" json:"req_limit_num"` // 请求次数限制
	LimitTime       int     `gorm:"column:limit_time" json:"limit_time"`       // 限制请求间隔时间
	Weight          float64 `json:"weight"`                                    // 权重
	Del             int     `json:"del"`                                       // 删除标志，0未删除1删除
	model.TimeColumns
}

// TableName 表名
func (f *FlowAdsRelModel) TableName() string {
	return "flow_ads_rel"
}

// GetsBy 查询
func (f *FlowAdsRelModel) GetsBy(query *tool.WhereQuery, order string) (list []*FlowAdsRelModel) {
	if order == "" {
		order = "update_time desc"
	}
	f.Table(f.TableName()).Where(query.QuerySql, query.QueryParams...).Order(order).Find(&list)
	return
}
