package model_api

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

type AdDeverPosModel struct {
	model.BaseModel `json:"-"`
	Id              int64   `gorm:"column:id" json:"id"`
	PosKeyType      string  `gorm:"column:pos_key_type" json:"pos_key_type"`       // 广告位类型
	AdSubType       int     `gorm:"column:ad_sub_type" json:"ad_sub_type"`         // 自定义广告的子类型 1（精品橱窗-焦点图）2（精品橱窗-应用墙）3（原生banner）
	Size            string  `gorm:"column:size" json:"size"`                       // 视频广告商编号
	DeverPosKey     string  `gorm:"column:dever_pos_key" json:"dever_pos_key"`     // 广告位KEY
	DeverPosName    string  `gorm:"column:dever_pos_name" json:"dever_pos_name"`   // 广告位名称
	State           int     `gorm:"column:state" json:"state"`                     // 1开启,0为关闭 默认为开启
	TestMode        int     `gorm:"column:test_mode" json:"test_mode"`             // 测试模式 0关闭1打开
	AppId           int64   `gorm:"column:app_id" json:"app_id"`                   // 应用appid
	DevId           int64   `gorm:"column:dev_id" json:"dev_id"`                   // 开发者id
	Rate            float64 `gorm:"column:rate" json:"rate"`                       // 广告位概率
	LimitNum        int     `gorm:"column:limit_num" json:"limit_num"`             // 限制次数
	PosDesc         string  `gorm:"column:pos_desc" json:"pos_desc"`               // 广告位描述
	AccountMethod   int     `gorm:"column:acounting_method" json:"account_method"` // 核算方式1:cpm  2:cpc
	Denominated     float64 `gorm:"column:denominated" json:"denominated"`         // 计价单位分
	Source          int     `gorm:"column:source" json:"source"`                   // 0为后台系统添加,1为网站开发者注册的应用
	CalState        int     `gorm:"column:cal_state" json:"cal_state"`             // 是否重新计算：0 否，1 是
	Del             int     `gorm:"column:del" json:"del"`                         // -1删除1是未删除
	model.TimeColumns
}

func AdDeverPosDB() *AdDeverPosModel {
	return &AdDeverPosModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

func (a *AdDeverPosModel) TableName() string {
	return "ad_dever_pos"
}

func (a *AdDeverPosModel) GetsBy(query *tool.WhereQuery, order string) (list []*AdDeverPosModel) {
	if order == "" {
		order = "update_time desc"
	}
	a.Where(query.QuerySql, query.QueryParams...).Order(order).Find(&list)
	return
}
