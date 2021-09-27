package model_api

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

type Channel struct {
	model.BaseModel `json:"-"`
	Id              int64  `json:"id" gorm:"column:id"`
	ChannelId       string `json:"channel_id" gorm:"column:channel_id"`
	ChannelName     string `json:"channel_name" gorm:"column:channel_name"`
	GroupId         int    `json:"group_id" gorm:"column:group_id"`
	IsCustom        int    `json:"is_custom" gorm:"column:is_custom"`             // 是非定制渠道
	IsCheckConfig   int    `json:"is_check_config" gorm:"column:is_check_config"` // 0:不检测1:检测
	AdsId           string `json:"ads_id" gorm:"column:ads_id"`                   // 定制渠道广告商
	DevId           int64  `json:"dev_id" gorm:"column:dev_id"`                   // 开发者id
	OperatorId      int64  `json:"operator_id" gorm:"column:operator_id"`
}

func ChannelDB() *Channel {
	return &Channel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

func (c *Channel) TableName() string {
	return "channel"
}

// GetChannel 传入条件获取渠道数据
func (c *Channel) GetChannel(query *tool.WhereQuery) (channel []*Channel) {
	c.Table(c.TableName()).Where(query.QuerySql, query.QueryParams...).Find(&channel)
	return
}
