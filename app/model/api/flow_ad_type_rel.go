package model_api

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

func FlowAdTypeRelDB() *FlowAdTypeRelModel {
	return &FlowAdTypeRelModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

type FlowAdTypeRelModel struct {
	model.BaseModel `json:"-"`
	Id              int64   `gorm:"column:id" json:"id"`
	FlowId          int64   `gorm:"column:flow_id" json:"flow_id"`                 // 流量配置id
	AppKey          string  `gorm:"column:app_key" json:"app_key"`                 //
	AdType          int     `gorm:"column:ad_type" json:"ad_type"`                 // 广告类型 1视频2插页
	Price           float64 `gorm:"column:price" json:"price"`                     // 底价
	Status          int     `gorm:"column:status" json:"status"`                   // 是否开启1是0否
	IsPriority      int     `gorm:"column:is_priority" json:"is_priority"`         // 是否使用优先1是0否
	IsUseDsp        int     `gorm:"column:is_use_dsp" json:"is_use_dsp"`           // 是否使用DSP 1是0否
	IsDelay         int     `gorm:"column:is_delay" json:"is_delay"`               // 尝鲜延迟加载1是0否
	Time            int64   `gorm:"column:time" json:"time"`                       // 尝鲜延迟加载时间,单位为秒
	Del             int     `gorm:"column:del" json:"del"`                         // 删除标志
	IsAppRel        int     `gorm:"column:is_app_rel" json:"is_app_rel"`           // 是否定向app配置0否1是
	IsBlockPolicy   int     `gorm:"column:is_block_policy" json:"is_block_policy"` // 是否定向广告位策略0否1是
	IsDefault       int     `gorm:"column:is_default" json:"is_default"`           // 是否使用默认配置
	model.TimeColumns
}

// TableName 表名
func (f *FlowAdTypeRelModel) TableName() string {
	return "flow_ad_type_rel"
}

func (f *FlowAdTypeRelModel) GetsBy(query *tool.WhereQuery, order string) (list []*FlowAdTypeRelModel) {
	if order == "" {
		order = "update_time desc"
	}
	f.Table(f.TableName()).Where(query.QuerySql, query.QueryParams...).Order(order).Find(&list)
	return
}
