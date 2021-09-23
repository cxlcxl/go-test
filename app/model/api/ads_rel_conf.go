package model_api

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

func AdsRelConfDB() *AdsRelConfModel {
	return &AdsRelConfModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

type AdsRelConfModel struct {
	model.BaseModel `json:"-"`
	Id              int            `gorm:"column:id" json:"id"` //
	AppKey          string         `gorm:"column:app_key" json:"app_key"`
	ConfName        string         `gorm:"column:conf_name" json:"conf_name"`     // 配置名称
	Version         int            `gorm:"column:version" json:"version"`         // 配置版本号，0为通用的全局配置
	Platform        int            `gorm:"column:platform" json:"platform"`       // 平台1安卓2ios
	Operator        string         `gorm:"column:operator" json:"operator"`       // 操作人
	UpdatedAt       tool.LocalTime `gorm:"column:update_time" json:"update_time"` // 限制次数
}

// TableName 表名
func (a *AdsRelConfModel) TableName() string {
	return "ads_rel_conf"
}

// GetsBy 数据查询
func (a *AdsRelConfModel) GetsBy(query *tool.WhereQuery) (adsRelConf []*AdsRelConfModel) {
	a.Where(query.QuerySql, query.QueryParams...).Order("version asc").Find(&adsRelConf)
	return
}
